- decode raw run:
  cat protoc-advance/decode-raw-option/proto/simple.bin | protoc --decode_raw
  result:
  1: 100
  2: 1
  3: "nanda"
  4: "\001\002\003\004\005"

- decode to proto message run
  cat protoc-advance/decode-raw-option/proto/simple.bin | protoc --decode=example.simple.Simple protoc-advance/decode-raw-option/proto/simple.proto
  result:
  id: 100
  is_simple: true
  name: "nanda"
  sample_lists: 1
  sample_lists: 2
  sample_lists: 3
  sample_lists: 4
  sample_lists: 5

  or you can add > "file_name" to decode to file

  cat protoc-advance/decode-raw-option/proto/simple.bin | protoc --decode=example.simple.Simple protoc-advance/decode-raw-option/proto/simple.proto > protoc-advance/decode-raw-option/proto/simple.txt

- encode
  cat protoc-advance/decode-raw-option/proto/simple.txt | protoc --encode=example.simple.Simple protoc-advance/decode-raw-option/proto/simple.proto > protoc-advance/decode-raw-option/proto/simple.pb
