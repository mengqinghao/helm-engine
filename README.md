# A [Helm3](https://github.com/helm/helm) Engine For Rest API With Go SDK

Helm3 摒弃了 Helm2 的 Tiller 架构，使用纯命令行的方式执行相关操作。如果想通过 Helm API 来实现相关功能，很遗憾官方并没有提供类似的服务。不过，因为官方提供了相对友好的 [Helm Go SDK](https://helm.sh/docs/topics/advanced/)，我们只需在此基础上做封装即可实现。为了不依赖cce集群的helm-wrapper插件，因此，Helm管理引擎提供平台统一管理helm的 RestAPI 供各服务调用

## Support API

* 如果某些API需要支持多个集群，则可以使用以下参数


| Params | Description |
| :- | :- |
| endpoint | 支持通过`endpoint`设置指定集群的APIServer地址 |
| token | 支持通过`token`设置指定集群的token认证  |
| kubeUserName | 支持通过`kubeUserName`设置指定集群的用户认证方式用户名  |
| kubePassword | 支持通过`kubePassword`设置指定集群的用户认证方式密码  |

helm 原生命令行和相关 API 对应关系：

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

> 此处 values 内容同 helm install `--values` 选项

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


> 此处 values 内容同 helm upgrade `--values` 选项

+ helm rollback
    - `PUT`
    - `/api/namespaces/:namespace/releases/:release/versions/:reversion`

PUT Body 可选:

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
| info | 支持 hooks/manifest/notes/values 信息，默认为 values |
| output | values 输出格式（仅当 info=values 时有效），支持 json/yaml，默认为 json |

+ helm release history
    - `GET`
    - `/api/namespaces/:namespace/releases/:release/histories`


+ helm show
    - `GET`
    - `/api/charts`

| Params | Description |
| :- | :- |
| chart  | 指定 chart 名，必填 |
| info   | 支持 all/readme/values/chart 信息，默认为 all |
| version | 支持版本指定，同命令行 |

+ helm search repo
    - `GET`
    - `/api/repositories/charts`
 
| Params | Description |
| :- | :- |
| keyword | 搜索关键字，必填 |
| version | 指定 chart version |
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
| chart | chart 包，必须为 .tgz 文件 |

+ list local charts
    - `GET`
    - `/api/charts/upload`

> 当前该版本处于 Alpha 状态，还没有经过大量的测试，只是把相关的功能测试了一遍，你也可以在此基础上自定义适合自身的版本。

### 响应

为了简化，所有请求统一返回 200 状态码，通过返回 Body 中的 Code 值来判断响应是否正常：

``` go
type respBody struct {
    Code  int         `json:"code"` // 0 or 1, 0 is ok, 1 is error
    Data  interface{} `json:"data,omitempty"`
    Error string      `json:"error,omitempty"`
}
```



