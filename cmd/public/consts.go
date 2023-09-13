package public

const IpfsClusterHost = "127.0.0.1:31333"
const IpfsHost = "127.0.0.1:30501"
const PinningHost = "pinning.solenopsys.org"

var (
	PATHS = map[string]string{
		"frontends":      "frontends/endpoints",
		"frontlibs":      "frontends/packages",
		"uimatrix":       "frontends/packages",
		"microfrontends": "frontends/modules",
		"tools":          "tools",
		"templates":      "templates",
		"backlibs":       "backends/libraries",
		"deployments":    "backends/deployments",
		"microservices":  "backends/modules",
	}

	PREFIXES = map[string]string{
		"fr": "frontends",
		"fl": "frontlibs",
		"ui": "uimatrix",
		"mf": "microfrontends",
		"tl": "tools",
		"tp": "templates",
		"bl": "backlibs",
		"dp": "deployments",
		"ms": "microservices",
	}
)
