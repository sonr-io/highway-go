module github.com/sonr-io/webauthn.io

go 1.17

replace github.com/sonr-io/sonr => ../sonr

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

require (
	github.com/cosmos/cosmos-sdk v0.45.1
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/duo-labs/webauthn v0.0.0-20220223184316-4d1cf2d34051
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/sessions v1.2.1
	github.com/jinzhu/gorm v1.9.16
	github.com/kataras/golog v0.1.7
	github.com/kataras/jwt v0.1.5
	github.com/sirupsen/logrus v1.8.1
	github.com/sonr-io/sonr v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.3.0
	github.com/spf13/viper v1.10.0
	github.com/stripe/stripe-go/v72 v72.93.0
	github.com/tendermint/starport v0.19.4
	go.buf.build/grpc/go/sonr-io/highway v1.2.24
	go.buf.build/grpc/go/sonr-io/sonr v1.2.14
	go.mongodb.org/mongo-driver v1.8.4
	google.golang.org/grpc v1.45.0
)

require (
	cloud.google.com/go v0.99.0 // indirect
	filippo.io/edwards25519 v1.0.0-beta.2 // indirect
	github.com/99designs/keyring v1.1.6 // indirect
	github.com/ChainSafe/go-schnorrkel v0.0.0-20200405005733-88cbf1b4c40d // indirect
	github.com/DataDog/zstd v1.4.5 // indirect
	github.com/armon/go-metrics v0.3.10 // indirect
	github.com/benbjohnson/clock v1.1.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/speakeasy v0.1.0 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/btcsuite/btcd v0.22.0-beta // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/census-instrumentation/opencensus-proto v0.3.0 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/cloudflare/cfssl v1.6.1 // indirect
	github.com/cncf/udpa/go v0.0.0-20210930031921-04548b0d99d4 // indirect
	github.com/cncf/xds/go v0.0.0-20211130200136-a8f946100490 // indirect
	github.com/confio/ics23/go v0.6.6 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/cosmos/btcutil v1.0.4 // indirect
	github.com/cosmos/go-bip39 v1.0.0 // indirect
	github.com/cosmos/iavl v0.17.3 // indirect
	github.com/cosmos/ledger-cosmos-go v0.11.1 // indirect
	github.com/cosmos/ledger-go v0.9.2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/danieljoos/wincred v1.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/davidlazar/go-crypto v0.0.0-20200604182044-b73af7476f6c // indirect
	github.com/dgraph-io/badger/v2 v2.2007.2 // indirect
	github.com/dgraph-io/ristretto v0.0.3 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/dustin/go-humanize v1.0.1-0.20200219035652-afde56e7acac // indirect
	github.com/dvsekhvalnov/jose2go v0.0.0-20200901110807-248326c1351b // indirect
	github.com/envoyproxy/go-control-plane v0.10.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.6.2 // indirect
	github.com/flynn/noise v1.0.0 // indirect
	github.com/form3tech-oss/jwt-go v3.2.3+incompatible // indirect
	github.com/francoispqt/gojay v1.2.13 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/fullstorydev/grpcurl v1.8.1 // indirect
	github.com/fxamacker/cbor/v2 v2.4.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/gogo/protobuf v1.3.3 // indirect
	github.com/golang-jwt/jwt/v4 v4.1.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/certificate-transparency-go v1.1.2-0.20210511102531-373a877eec92 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/gtank/merlin v0.1.1 // indirect
	github.com/gtank/ristretto255 v0.1.2 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hdevalence/ed25519consensus v0.0.0-20210204194344-59a8610d2b87 // indirect
	github.com/huin/goupnp v1.0.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/ipfs/go-cid v0.1.0 // indirect
	github.com/ipfs/go-datastore v0.5.1 // indirect
	github.com/ipfs/go-ipfs-util v0.0.2 // indirect
	github.com/ipfs/go-ipns v0.1.2 // indirect
	github.com/ipfs/go-log v1.0.5 // indirect
	github.com/ipfs/go-log/v2 v2.3.0 // indirect
	github.com/ipld/go-ipld-prime v0.14.2 // indirect
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/jbenet/go-temp-err-catcher v0.1.0 // indirect
	github.com/jbenet/goprocess v0.1.4 // indirect
	github.com/jhump/protoreflect v1.9.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kataras/pio v0.0.10 // indirect
	github.com/keybase/go-keychain v0.0.0-20190712205309-48d3d31d256d // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/koron/go-ssdp v0.0.2 // indirect
	github.com/libp2p/go-addr-util v0.1.0 // indirect
	github.com/libp2p/go-buffer-pool v0.0.2 // indirect
	github.com/libp2p/go-cidranger v1.1.0 // indirect
	github.com/libp2p/go-conn-security-multistream v0.3.0 // indirect
	github.com/libp2p/go-eventbus v0.2.1 // indirect
	github.com/libp2p/go-flow-metrics v0.0.3 // indirect
	github.com/libp2p/go-libp2p v0.16.0 // indirect
	github.com/libp2p/go-libp2p-asn-util v0.1.0 // indirect
	github.com/libp2p/go-libp2p-autonat v0.6.0 // indirect
	github.com/libp2p/go-libp2p-blankhost v0.2.0 // indirect
	github.com/libp2p/go-libp2p-connmgr v0.2.4 // indirect
	github.com/libp2p/go-libp2p-core v0.11.0 // indirect
	github.com/libp2p/go-libp2p-discovery v0.6.0 // indirect
	github.com/libp2p/go-libp2p-kad-dht v0.15.0 // indirect
	github.com/libp2p/go-libp2p-kbucket v0.4.7 // indirect
	github.com/libp2p/go-libp2p-mplex v0.4.1 // indirect
	github.com/libp2p/go-libp2p-nat v0.1.0 // indirect
	github.com/libp2p/go-libp2p-noise v0.3.0 // indirect
	github.com/libp2p/go-libp2p-peerstore v0.4.0 // indirect
	github.com/libp2p/go-libp2p-pnet v0.2.0 // indirect
	github.com/libp2p/go-libp2p-pubsub v0.6.0 // indirect
	github.com/libp2p/go-libp2p-quic-transport v0.15.0 // indirect
	github.com/libp2p/go-libp2p-record v0.1.3 // indirect
	github.com/libp2p/go-libp2p-swarm v0.8.0 // indirect
	github.com/libp2p/go-libp2p-tls v0.3.1 // indirect
	github.com/libp2p/go-libp2p-transport-upgrader v0.5.0 // indirect
	github.com/libp2p/go-libp2p-yamux v0.6.0 // indirect
	github.com/libp2p/go-maddr-filter v0.1.0 // indirect
	github.com/libp2p/go-mplex v0.3.0 // indirect
	github.com/libp2p/go-msgio v0.1.0 // indirect
	github.com/libp2p/go-nat v0.1.0 // indirect
	github.com/libp2p/go-netroute v0.1.6 // indirect
	github.com/libp2p/go-openssl v0.0.7 // indirect
	github.com/libp2p/go-reuseport v0.1.0 // indirect
	github.com/libp2p/go-reuseport-transport v0.1.0 // indirect
	github.com/libp2p/go-sockaddr v0.1.1 // indirect
	github.com/libp2p/go-stream-muxer-multistream v0.3.0 // indirect
	github.com/libp2p/go-tcp-transport v0.4.0 // indirect
	github.com/libp2p/go-ws-transport v0.5.0 // indirect
	github.com/libp2p/go-yamux/v2 v2.3.0 // indirect
	github.com/libp2p/zeroconf/v2 v2.1.1 // indirect
	github.com/lucas-clemente/quic-go v0.24.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/marten-seemann/qtls-go1-16 v0.1.4 // indirect
	github.com/marten-seemann/qtls-go1-17 v0.1.0 // indirect
	github.com/marten-seemann/tcp v0.0.0-20210406111302-dfbc87cc63fd // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.12 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/mikioh/tcpinfo v0.0.0-20190314235526-30a79bb1804b // indirect
	github.com/mikioh/tcpopt v0.0.0-20190314235656-172688c1accc // indirect
	github.com/mimoo/StrobeGo v0.0.0-20181016162300-f8f6d4d2b643 // indirect
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/multiformats/go-base32 v0.0.3 // indirect
	github.com/multiformats/go-base36 v0.1.0 // indirect
	github.com/multiformats/go-multiaddr v0.4.1 // indirect
	github.com/multiformats/go-multiaddr-dns v0.3.1 // indirect
	github.com/multiformats/go-multiaddr-fmt v0.1.0 // indirect
	github.com/multiformats/go-multibase v0.0.3 // indirect
	github.com/multiformats/go-multicodec v0.3.0 // indirect
	github.com/multiformats/go-multihash v0.1.0 // indirect
	github.com/multiformats/go-multistream v0.2.2 // indirect
	github.com/multiformats/go-varint v0.0.6 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/polydawn/refmt v0.0.0-20201211092308-30ac6d18308e // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.30.0 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0 // indirect
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rs/cors v1.7.0 // indirect
	github.com/rs/zerolog v1.26.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sasha-s/go-deadlock v0.2.1-0.20190427202633-1595213edefa // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/spacemonkeygo/spacelog v0.0.0-20180420211403-2296661a0572 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20200815110645-5c35d600f0ca // indirect
	github.com/tecbot/gorocksdb v0.0.0-20191217155057-f0fad39f321c // indirect
	github.com/tendermint/btcd v0.1.1 // indirect
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15 // indirect
	github.com/tendermint/go-amino v0.16.0 // indirect
	github.com/tendermint/tendermint v0.34.14 // indirect
	github.com/tendermint/tm-db v0.6.4 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20201229170055-e5319fda7802 // indirect
	github.com/urfave/cli v1.22.5 // indirect
	github.com/whyrusleeping/go-keyspace v0.0.0-20160322163242-5b898ac5add1 // indirect
	github.com/whyrusleeping/multiaddr-filter v0.0.0-20160516205228-e903e4adabd7 // indirect
	github.com/whyrusleeping/timecache v0.0.0-20160911033111-cfcb2f1abfee // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	github.com/zondax/hid v0.9.0 // indirect
	go.buf.build/grpc/go/cosmos/cosmos-sdk v1.2.31 // indirect
	go.buf.build/grpc/go/cosmos/gogo-proto v1.2.1 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	go.etcd.io/etcd/api/v3 v3.5.1 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.1 // indirect
	go.etcd.io/etcd/client/v2 v2.305.1 // indirect
	go.etcd.io/etcd/client/v3 v3.5.0-alpha.0 // indirect
	go.etcd.io/etcd/etcdctl/v3 v3.5.0-alpha.0 // indirect
	go.etcd.io/etcd/pkg/v3 v3.5.0-alpha.0 // indirect
	go.etcd.io/etcd/raft/v3 v3.5.0-alpha.0 // indirect
	go.etcd.io/etcd/server/v3 v3.5.0-alpha.0 // indirect
	go.etcd.io/etcd/tests/v3 v3.5.0-alpha.0 // indirect
	go.etcd.io/etcd/v3 v3.5.0-alpha.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/crypto v0.0.0-20211215165025-cf75a172585e // indirect
	golang.org/x/mod v0.5.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/oauth2 v0.0.0-20220309155454-6242fa91716a // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	golang.org/x/tools v0.1.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220302033224-9aa15565e42a // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/cheggaaa/pb.v1 v1.0.28 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	lukechampine.com/blake3 v1.1.6 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
