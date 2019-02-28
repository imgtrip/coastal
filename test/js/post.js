let meta = cm({token: "TestToken0"});
client.showPost({id: 1}, meta, pr)

// grpcc.js --insecure --proto ../../protobuf/imgtrip.proto --address localhost:50050 --exec ./post.js

// grpcc -i -p ../../protobuf/imgtrip.proto -a localhost:50050 --exec post.js