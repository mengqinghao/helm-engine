# A [Helm3](https://github.com/helm/helm) Engine For Rest API With [Go SDK](https://helm.sh/docs/topics/advanced/#go-sdk)

+ [中文文档](README.md)

helm-engine is a helm3 management engine for rest API with [helm Go SDK](https://helm.sh/docs/topics/advanced/#go-sdk). With helm-engine, you can use HTTP RESTFul API do something like helm commondline (install/uninstall/upgrade/get/list/rollback...).

## Support API


* If there are some APIs need to support Specify clusters, you can use the parameters below
 (default cluster without the parameters below)


| Params | Description |
| :- | :- |
| endpoint | Support Specify clusters APIServer by the`endpoint`  |
| token | Support Specify clusters auth by the`token`  |
| kubeUserName | Support Specify clusters auth by the`kubeUserName`  |
| kubePassword | Support Specify clusters auth by the`kubePassword`  |



+ helm install
    - `POST`
    - `/api/namespaces/:namespace/releases/:release?chart=<chartName>`

POST Body: 

``` json
{
    "dry_run": false,           // `--dry-run`
    "disable_hooks": false,     // `--no-hooks`
    "wait": false,              // `--wait`
    "devel": false,             // `--false`
    "description": "",          // `--description`
    "atomic": false,            // `--atomic`
    "skip_crds": false,         // `--skip-crds`
    "sub_notes": false,         // `--render-subchart-notes`
    "create_namespace": false,  // `--create-namespace`
    "dependency_update": false, // `--dependency-update`
    "values": "",               // `--values`
    "set": [],                  // `--set`
    "set_string": [],           // `--set-string`
    "ca_file": "",              // `--ca-file`
    "cert_file": "",            // `--cert-file`
    "key_file": "",             // `--key-file`
    "insecure_skip_verify": "", // `--insecure-skip-verify`
    "keyring": "",              // `--keyring`
    "password": "",             // `--password`
    "repo": "",                 // `--repo`
    "username": "",             // `--username`
    "verify": false,            // `--verify`
    "version": ""               // `--version`
}
```

> `"values"` -> helm install `--values` option 

+ helm uninstall
    - `DELETE`
    - `/api/namespaces/:namespace/releases/:release`


+ helm upgrade
    - `PUT`
    - `/api/namespaces/:namespace/releases/:release?chart=<chartName>`

PUT Body: 

``` json
{
    "dry_run": false,           // `--dry-run`
    "disable_hooks": false,     // `--no-hooks`
    "wait": false,              // `--wait`
    "devel": false,             // `--false`
    "description": "",          // `--description`
    "atomic": false,            // `--atomic`
    "skip_crds": false,         // `--skip-crds`
    "sub_notes": false,         // `--render-subchart-notes`
    "force": false,             // `--force`
    "install": false,           // `--install`
    "recreate": false,          // `--recreate`
    "cleanup_on_fail": false,   // `--cleanup-on-fail`
    "values": "",               // `--values`
    "set": [],                  // `--set`
    "set_string": [],           // `--set-string`
    "ca_file": "",              // `--ca-file`
    "cert_file": "",            // `--cert-file`
    "key_file": "",             // `--key-file`
    "insecure_skip_verify": "", // `--insecure-skip-verify`
    "keyring": "",              // `--keyring`
    "password": "",             // `--password`
    "repo": "",                 // `--repo`
    "username": "",             // `--username`
    "verify": false,            // `--verify`
    "version": ""               // `--version`
}
```

> `"values"` -> helm install `--values` option 

+ helm rollback
    - `PUT`
    - `/api/namespaces/:namespace/releases/:release/versions/:reversion`

PUT Body (optional):

``` json
{
    "dry_run": false,           // `--dry-run`
    "disable_hooks": false,     // `--no-hooks`
    "wait": false,              // `--wait`
    "force": false,             // `--force`
    "recreate": false,          // `--recreate`
    "cleanup_on_fail": false,   // `--cleanup-on-fail`
    "history_max":              // `--history-max` int
}
```

+ helm list
    - `GET`
    - `/api/namespaces/:namespace/releases`

Body:

``` json
{
    "all": false,               // `--all`
    "all_namespaces": false,    // `--all-namespaces`
    "by_date": false,           // `--date`
    "sort_reverse": false,      // `--reverse`
    "limit":  ,                 // `--max`
    "offset": ,                 // `--offset`
    "filter": "",               // `--filter`
    "uninstalled": false,       // `--uninstalled`
    "uninstalling": false,      // `--uninstalling`
    "superseded": false,        // `--superseded`
    "failed": false,            // `--failed`
    "deployed": false,          // `--deployed`
    "pending": false            // `--pending`
}
```

+ helm get
    - `GET`
    - `/api/namespaces/:namespace/releases/:release`

| Params | Description |
| :- | :- |
| info | support hooks/manifest/notes/values, default values |
| output | get values output format (only info==values), support json/yaml, default json |

+ helm release history
    - `GET`
    - `/api/namespaces/:namespace/releases/:release/histories`


+ helm show
    - `GET`
    - `/api/charts`

| Params | Description |
| :- | :- |
| chart  | chart name, required|
| info   | support all/readme/values/chart, default all |
| version | --version |

+ helm search repo
    - `GET`
    - `/api/repositories/charts`

| Params | Description |
| :- | :- |
| keyword | search keyword，required |
| version | chart version |
| versions | if "true", all versions |

+ helm repo update
    - `PUT`
    - `/api/repositories`

+ helm env
    - `GET`
    - `/api/envs`

+ upload chart
    - `POST`
    - `/api/charts/upload`

 
| Params | Description |
| :- | :- |
| chart | upload chart file, with suffix .tgz |

+ list local charts
    - `GET`
    - `/api/charts/upload`

> __Notes:__ helm-engine is Alpha status, no more test

### Response 


``` go
type respBody struct {
    Code  int         `json:"code"` // 0 or 1, 0 is ok, 1 is error
    Data  interface{} `json:"data,omitempty"`
    Error string      `json:"error,omitempty"`
}
```



