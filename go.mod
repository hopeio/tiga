module github.com/hopeio/lemon

go 1.21

toolchain go1.21.0

require (
	bou.ke/monkey v1.0.2
	github.com/99designs/gqlgen v0.17.36
	github.com/IBM/sarama v1.41.0
	github.com/alta/protopatch v0.5.3
	github.com/aws/aws-sdk-go v1.44.332
	github.com/boombuler/barcode v1.0.1
	github.com/cespare/xxhash v1.1.0
	github.com/cockroachdb/pebble v0.0.0-20230826001808-0b401ee526b8
	github.com/dgraph-io/badger/v3 v3.2103.5
	github.com/dgraph-io/ristretto v0.1.1
	github.com/fsnotify/fsnotify v1.6.0
	github.com/gin-gonic/gin v1.9.1
	github.com/go-oauth2/oauth2/v4 v4.5.2
	github.com/go-openapi/loads v0.21.2
	github.com/go-openapi/runtime v0.26.0
	github.com/go-openapi/spec v0.20.9
	github.com/go-openapi/swag v0.22.4
	github.com/go-playground/locales v0.14.1
	github.com/go-playground/universal-translator v0.18.1
	github.com/go-playground/validator/v10 v10.15.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/gofiber/fiber/v2 v2.49.0
	github.com/google/flatbuffers v23.5.26+incompatible // indirect
	github.com/google/uuid v1.3.1
	github.com/gorilla/sessions v1.2.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.17.1
	github.com/hashicorp/go-metrics v0.5.1
	github.com/json-iterator/go v1.1.12
	github.com/markbates/goth v1.77.0
	github.com/microcosm-cc/bluemonday v1.0.25
	github.com/modern-go/reflect2 v1.0.2
	github.com/mozillazg/go-pinyin v0.20.0
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/nsqio/go-nsq v1.1.0
	github.com/olivere/elastic/v7 v7.0.32
	github.com/pelletier/go-toml v1.9.5
	github.com/prometheus/client_golang v1.16.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/russross/blackfriday v1.6.0
	github.com/spf13/viper v1.16.0
	github.com/stretchr/testify v1.8.4
	github.com/tealeg/xlsx/v3 v3.3.0
	github.com/tidwall/pretty v1.2.1
	github.com/ugorji/go/codec v1.2.11
	github.com/valyala/fasthttp v1.48.0
	github.com/xuri/excelize/v2 v2.8.0
	go.etcd.io/etcd/client/v3 v3.5.9
	go.opencensus.io v0.24.0
	go.uber.org/atomic v1.11.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.25.0
	golang.org/x/net v0.14.0
	golang.org/x/oauth2 v0.11.0
	golang.org/x/tools v0.12.1-0.20230815132531-74c255bcf846
	google.golang.org/genproto v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/grpc v1.57.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.3.0
	google.golang.org/protobuf v1.31.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/mysql v1.5.1
	gorm.io/driver/postgres v1.5.2
	gorm.io/driver/sqlite v1.5.3
	gorm.io/gorm v1.25.4
	gorm.io/plugin/prometheus v0.0.0-20230504115745-1aec2356381b
)

require (
	github.com/blastrain/vitess-sqlparser v0.0.0-20201030050434-a139afbb1aba
	github.com/bytedance/sonic v1.10.0
	github.com/cenkalti/backoff/v4 v4.2.1
	github.com/danielvladco/go-proto-gql v0.10.1-0.20221227181908-22fca0a9469c
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc
	github.com/elastic/go-elasticsearch/v8 v8.9.0
	github.com/golang-jwt/jwt/v5 v5.0.0
	github.com/jhump/protoreflect v1.15.2
	github.com/lni/dragonboat/v4 v4.0.0-20230709075559-54497b9553be
	github.com/lni/goutils v1.3.1-0.20220604063047-388d67b4dbc4
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f
	github.com/mwitkow/grpc-proxy v0.0.0-20230212185441-f345521cb9c9
	github.com/nacos-group/nacos-sdk-go/v2 v2.2.3
	github.com/quic-go/quic-go v0.38.1
	github.com/rs/cors v1.9.0
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	github.com/urfave/cli/v2 v2.25.7
	github.com/vektah/gqlparser/v2 v2.5.8
	golang.org/x/exp v0.0.0-20230817173708-d852ddb80c63
	golang.org/x/image v0.11.0
	golang.org/x/text v0.12.0
	golang.org/x/time v0.3.0
	google.golang.org/genproto/googleapis/api v0.0.0-20230822172742-b8732ec3820d
	nhooyr.io/websocket v1.8.7
)

require (
	cloud.google.com/go v0.110.7 // indirect
	cloud.google.com/go/compute v1.23.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/firestore v1.12.0 // indirect
	cloud.google.com/go/longrunning v0.5.1 // indirect
	github.com/DataDog/zstd v1.5.5 // indirect
	github.com/VictoriaMetrics/metrics v1.24.0 // indirect
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.62.524 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bufbuild/protocompile v0.6.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/cockroachdb/errors v1.11.0 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/cockroachdb/tokenbucket v0.0.0-20230807174530-cc333fc44b06 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eapache/go-resiliency v1.4.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/elastic/elastic-transport-go/v8 v8.3.0 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/frankban/quicktest v1.14.6 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/getsentry/sentry-go v0.23.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/analysis v0.21.4 // indirect
	github.com/go-openapi/errors v0.20.4 // indirect
	github.com/go-openapi/jsonpointer v0.20.0 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/strfmt v0.21.7 // indirect
	github.com/go-openapi/validate v0.22.1 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.1.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20230821062121-407c9e7a662f // indirect
	github.com/google/s2a-go v0.1.5 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.5 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/consul/api v1.24.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.6 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/memberlist v0.5.0 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/juju/errors v1.0.0 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/lni/vfs v0.2.1-0.20220616104132-8852fd867376 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/markbates/going v1.0.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/miekg/dns v1.1.55 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/mrjones/oauth v0.0.0-20190623134757-126b35219450 // indirect
	github.com/nats-io/jwt/v2 v2.5.0 // indirect
	github.com/nats-io/nats.go v1.28.0 // indirect
	github.com/nats-io/nkeys v0.4.4 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/onsi/ginkgo/v2 v2.12.0 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20220228012449-10b1cf09e00b // indirect
	github.com/pelletier/go-toml/v2 v2.0.9 // indirect
	github.com/peterbourgon/diskv/v3 v3.0.1 // indirect
	github.com/pierrec/lz4/v4 v4.1.18 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.11.1 // indirect
	github.com/quic-go/qpack v0.4.0 // indirect
	github.com/quic-go/qtls-go1-20 v0.3.3 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/richardlehane/mscfb v1.0.4 // indirect
	github.com/richardlehane/msoleps v1.0.3 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/rogpeppe/fastuuid v1.2.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sagikazarmark/crypt v0.13.0 // indirect
	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529 // indirect
	github.com/shabbyrobe/xmlwriter v0.0.0-20230525083848-85336ec334fa // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fastrand v1.1.0 // indirect
	github.com/valyala/histogram v1.2.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	github.com/xuri/efp v0.0.0-20230802181842-ad255f2331ca // indirect
	github.com/xuri/nfp v0.0.0-20230819163627-dc951e3ffe1a // indirect
	go.etcd.io/etcd/api/v3 v3.5.9 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.9 // indirect
	go.etcd.io/etcd/client/v2 v2.305.9 // indirect
	go.mongodb.org/mongo-driver v1.12.1 // indirect
	golang.org/x/arch v0.4.0 // indirect
	golang.org/x/crypto v0.12.0 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.138.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
)

replace github.com/armon/go-metrics v0.4.1 => github.com/hashicorp/go-metrics v0.5.1
