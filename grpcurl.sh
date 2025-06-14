# list
grpcurl -plaintext \
  -import-path ./proto \
  -import-path /opt/homebrew/include \
  -proto service_bank_bhutan.proto \
  localhost:9000 \
  list pb.BankBhutan

# create user
grpcurl -plaintext \
  -import-path ./proto \
  -import-path /opt/homebrew/include \
  -proto service_bank_bhutan.proto \
  -d '{"username":"jigme", "full_name":"Jigme Tenzin", "email":"jigme@example.com", "password":"secure123"}' \
  localhost:9000 \
  pb.BankBhutan/CreateUser

# login user
grpcurl -plaintext \
  -import-path ./proto \
  -import-path /opt/homebrew/include \
  -proto service_bank_bhutan.proto \
  -d '{"username":"jigme", "password":"secure123"}' \
  localhost:9000 \
  pb.BankBhutan/LoginUser
