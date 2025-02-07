package main

import (
	"flag"
	"fmt"
	"github.com/jtacoma/uritemplates"
	"github.com/whosonfirst/go-rasterzen/http"
	"github.com/whosonfirst/go-rasterzen/nextzen"
	"github.com/whosonfirst/go-rasterzen/server"
	"github.com/whosonfirst/go-rasterzen/tile"
	"github.com/whosonfirst/go-whosonfirst-cache"
	"github.com/whosonfirst/go-whosonfirst-cache-s3"
	"github.com/whosonfirst/go-whosonfirst-cli/flags"
	"log"
	gohttp "net/http"
	gourl "net/url"
	"os"
	"strings"
)

func main() {

	config := flag.String("config", "", "Read some or all flags from an ini-style config file. Values in the config file take precedence over command line flags.")
	section := flag.String("section", "rasterd", "A valid ini-style config file section.")

	var proto = flag.String("protocol", "http", "The protocol for wof-staticd server to listen on. Valid protocols are: http, lambda.")
	var host = flag.String("host", "localhost", "The host for rasterd to listen for requests on.")
	var port = flag.Int("port", 8080, "The port for rasterd to listen for requests on.")

	do_www := flag.Bool("www", false, "Enable a simple web interface with a slippy map (at /) for testing and debugging.")
	www_debug := flag.Bool("www-debug", false, "Enable debugging features for the web interface.")
	www_tile_format := flag.String("www-tile-format", "svg", "Valid options are: png, svg.")

	no_cache := flag.Bool("no-cache", false, "Disable all caching.")
	go_cache := flag.Bool("go-cache", false, "Cache tiles with an in-memory (go-cache) cache.")
	fs_cache := flag.Bool("fs-cache", false, "Cache tiles with a filesystem-based cache.")
	fs_root := flag.String("fs-root", "", "The root of your filesystem cache. If empty rasterd will try to use the current working directory.")
	s3_cache := flag.Bool("s3-cache", false, "Cache tiles with a S3-based cache.")
	s3_dsn := flag.String("s3-dsn", "", "A valid go-whosonfirst-aws DSN string")
	s3_opts := flag.String("s3-opts", "", "A valid go-whosonfirst-cache-s3 options string")

	nextzen_apikey := flag.String("nextzen-apikey", "", "A valid Nextzen API key.")
	nextzen_origin := flag.String("nextzen-origin", "", "An optional HTTP 'Origin' host to pass along with your Nextzen requests.")
	nextzen_debug := flag.Bool("nextzen-debug", false, "Log requests (to STDOUT) to Nextzen tile servers.")
	nextzen_uri := flag.String("nextzen-uri", "", "A valid URI template (RFC 6570) pointing to a custom Nextzen endpoint.")

	png_handler := flag.Bool("png-handler", true, "Enable the PNG tile handler.")
	svg_handler := flag.Bool("svg-handler", true, "Enable the SVG tile handler.")
	rasterzen_handler := flag.Bool("rasterzen-handler", false, "Enable the Rasterzen tile handler.")
	geojson_handler := flag.Bool("geojson-handler", false, "Enable the GeoJSON tile handler.")

	custom_svg_options := flag.String("svg-options", "", "Custom RasterzenSVGOptions data. This may be a path to a JSON config file or a valid JSON string.")

	var path_png = flag.String("path-png", "/png/", "The path that PNG tiles should be served from")
	var path_svg = flag.String("path-svg", "/svg/", "The path that SVG tiles should be served from")
	var path_geojson = flag.String("path-geojson", "/geojson/", "The path that GeoJSON tiles should be served from")
	var path_rasterzen = flag.String("path-rasterzen", "/rasterzen/", "The path that Rasterzen tiles should be served from")

	flag.Parse()

	if *nextzen_apikey == "" && *nextzen_uri == "" {
		log.Println("Missing -nextzen-apikey flag. Unless you've already cached your tiles you won't be able to fetch tiles to render!")
	}

	if *config != "" {

		err := flags.SetFlagsFromConfig(*config, *section)

		if err != nil {
			log.Fatal(err)
		}

	} else {

		err := flags.SetFlagsFromEnvVars("RASTERD")

		if err != nil {
			log.Fatal(err)
		}
	}

	if *no_cache {

		log.Println("disable all cache layers")

		*go_cache = false
		*fs_cache = false
	}

	nz_opts := &nextzen.Options{
		ApiKey: *nextzen_apikey,
		Origin: *nextzen_origin,
		Debug:  *nextzen_debug,
	}

	if *nextzen_uri != "" {

		template, err := uritemplates.Parse(*nextzen_uri)

		if err != nil {
			log.Fatal(err)
		}

		nz_opts.URITemplate = template
	}

	caches := make([]cache.Cache, 0)

	if *go_cache {

		log.Println("enable go-cache cache layer")

		opts, err := cache.DefaultGoCacheOptions()

		if err != nil {
			log.Fatal(err)
		}

		c, err := cache.NewGoCache(opts)

		if err != nil {
			log.Fatal(err)
		}

		caches = append(caches, c)
	}

	if *fs_cache {

		log.Println("enable filesystem cache layer")

		if *fs_root == "" {

			cwd, err := os.Getwd()

			if err != nil {
				log.Fatal(err)
			}

			*fs_root = cwd
		}

		c, err := cache.NewFSCache(*fs_root)

		if err != nil {
			log.Fatal(err)
		}

		caches = append(caches, c)
	}

	if *s3_cache {

		log.Println("enable S3 cache layer")

		opts, err := s3.NewS3CacheOptionsFromString(*s3_opts)

		if err != nil {
			log.Fatal(err)
		}

		c, err := s3.NewS3Cache(*s3_dsn, opts)

		if err != nil {
			log.Fatal(err)
		}

		caches = append(caches, c)
	}

	if len(caches) == 0 {

		// because we still need to pass a cache.Cache thingy
		// around (20180612/thisisaaronland)

		c, err := cache.NewNullCache()

		if err != nil {
			log.Fatal(err)
		}

		caches = append(caches, c)
	}

	c, err := cache.NewMultiCache(caches)

	if err != nil {
		log.Fatal(err)
	}

	mux := gohttp.NewServeMux()

	if *png_handler || *svg_handler {

		var svg_opts *tile.RasterzenSVGOptions

		if *custom_svg_options != "" {

			var opts *tile.RasterzenSVGOptions

			if strings.HasPrefix(*custom_svg_options, "{") {
				opts, err = tile.RasterzenSVGOptionsFromString(*custom_svg_options)
			} else {
				opts, err = tile.RasterzenSVGOptionsFromFile(*custom_svg_options)
			}

			if err != nil {
				log.Fatal(err)
			}

			svg_opts = opts

		} else {

			opts, err := tile.DefaultRasterzenSVGOptions()

			if err != nil {
				log.Fatal(err)
			}

			svg_opts = opts
		}

		if *png_handler {

			log.Println("enable PNG handler")

			h, err := http.NewPNGHandler(c, nz_opts, svg_opts)

			if err != nil {
				log.Fatal(err)
			}

			mux.Handle(*path_png, h)
		}

		if *svg_handler {

			log.Println("enable SVG handler")

			h, err := http.NewSVGHandler(c, nz_opts, svg_opts)

			if err != nil {
				log.Fatal(err)
			}

			mux.Handle(*path_svg, h)
		}
	}

	if *geojson_handler {

		log.Println("enable GeoJSON handler")

		h, err := http.NewGeoJSONHandler(c, nz_opts)

		if err != nil {
			log.Fatal(err)
		}

		mux.Handle(*path_geojson, h)
	}

	if *rasterzen_handler {

		log.Println("enable Rasterzen handler")

		h, err := http.NewRasterzenHandler(c, nz_opts)

		if err != nil {
			log.Fatal(err)
		}

		mux.Handle(*path_rasterzen, h)
	}

	if *do_www {

		// We have (need) two separate handlers (and by extension bundled assets
		// in ur go binary tools) for the www handler which is not great but that's
		// it goes sometimes. The first is a standard "static" assets (js, css)
		// wrapper than unfortunately needs to be stored in http/assetsfs.go because
		// of namespaces and private function names. The second is a bundled Go
		// template (index.html) specifically so we can assign a nextzen API key.
		// I suppose we could have also written a middleware handler to modify
		// something coming out of the (first) static asset bundle but I'm not convinced
		// that isn't more confusing that this set up. The static bundle depends
		// on 'go-bindata-assetfs' and the template bundle uses 'go-bindata-html-template'.
		// As the names suggest everything uses 'go-bindata'. If you make changes to any
		// of the static assets or the templates you'll need to rebuild them using the
		// handy 'make assets bin' or 'make rebuild' Makefile targets. Good times...
		// (20181102/thisisaaronland)

		if *nextzen_apikey == "" {
			log.Fatal("You must pass a -nextzen-apikey parameter for the local www server to work")
		}

		log.Println("enable WWW handler")

		static_h, err := http.StaticHandler()

		if err != nil {
			log.Fatal(err)
		}

		www_opts := &http.WWWHandlerOptions{
			NextzenAPIKey: *nextzen_apikey,
			Debug:         *www_debug,
			TileFormat:    *www_tile_format,
		}

		www_h, err := http.WWWHandler(www_opts)

		if err != nil {
			log.Fatal(err)
		}

		mux.Handle("/javascript/", static_h)
		mux.Handle("/css/", static_h)
		mux.Handle("/", www_h)
	}

	address := fmt.Sprintf("http://%s:%d", *host, *port)

	u, err := gourl.Parse(address)

	if err != nil {
		log.Fatal(err)
	}

	s, err := server.NewStaticServer(*proto, u)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on %s\n", s.Address())

	err = s.ListenAndServe(mux)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
