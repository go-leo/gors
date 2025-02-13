protoc \
--proto_path=. \
--proto_path=../third_party \
--proto_path=../../ \
--go_out=. \
--go_opt=paths=source_relative \
--gors-gorilla_out=. \
--gors-gorilla_opt=paths=source_relative \
*/*.proto
