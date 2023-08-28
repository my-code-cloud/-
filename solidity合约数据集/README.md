# Solidity智能合约数据集

* 合约来源：Smart Contract Weakness Classification (SWC)

|合约|下载地址|漏洞|
|:--|:--|:--|
|arbitrary_location_write_simple_fixed.sol|https://swcregistry.io/docs/SWC-124/|controlled-array-length|
|arbitrary_location_write_simple.sol|https://swcregistry.io/docs/SWC-124/|controlled-array-length|
|crypto_roulette_fixed.sol|https://swcregistry.io/docs/SWC-109/|weak-prng, controlled-array-length, incorrect-equality（误报）, uninitialized-local|
|crypto_roulette.sol|https://swcregistry.io/docs/SWC-109/|weak-prng, controlled-array-length, incorrect-equality（误报）， uninitialized-storage|
|deposit_box_fixed.sol|https://swcregistry.io/docs/SWC-135/|uninitialized-state（漏报）|
|deposit_box.sol|https://swcregistry.io/docs/SWC-135/|uninitialized-state|
|deprecated_simple_fixed.sol|https://swcregistry.io/docs/SWC-111/|suicidal, unchecked-lowlevel|
|deprecated_simple.sol|https://swcregistry.io/docs/SWC-111/|suicidal, unchecked-lowlevel|
|dos_address.sol|https://swcregistry.io/docs/SWC-128/|controlled-array-length|
|FunctionTypes.sol|https://swcregistry.io/docs/SWC-127/|arbitrary-send|
|guess_the_random_number_fixed.sol|https://swcregistry.io/docs/SWC-120/|2 incorrect-equality（误报）|
|guess_the_random_number.sol|https://swcregistry.io/docs/SWC-120/|arbitrary-send, 2 incorrect-equality（误报）|
|mapping_perfomance_2.sol|https://swcregistry.io/docs/SWC-110/|uninitialized-state|
|mapping_performance_1.sol|https://swcregistry.io/docs/SWC-110/|uninitialized-state|
|mapping_write.sol|https://swcregistry.io/docs/SWC-124/|controlled-array-length, uninitialized-state|
|mycontract.sol|https://swcregistry.io/docs/SWC-115/|arbitrary-send, tx-origin|
|odd_even_fixed.sol|https://swcregistry.io/docs/SWC-136/|reentrancy-eth, 2 uninitialized-local, 2 unchecked-lowlevel|
|odd_even.sol|https://swcregistry.io/docs/SWC-136/|reentrancy-eth|
|proxy_fixed.sol|https://swcregistry.io/docs/SWC-112/|controlled-delegatecall|
|proxy_pattern_false_positive.sol|https://swcregistry.io/docs/SWC-112/|controlled-delegatecall, uninitialized-state|
|proxy.sol|https://swcregistry.io/docs/SWC-112/|controlled-delegatecall|
|random_number_generator.sol|https://swcregistry.io/docs/SWC-120/|weak-prng|
|relayer_fixed.sol|https://swcregistry.io/docs/SWC-126/|unchecked-lowlevel|
|rubixi.sol|https://swcregistry.io/docs/SWC-105/|controlled-array-length, 4 unchecked-send, divide-before-multiply|
|send_loop.sol|http://39.103.152.161/|uninitialized-state, uninitialized-local|
|sha_of_sha_2_mappings.sol|https://swcregistry.io/docs/SWC-110/|uninitialized-state|
|sha_of_sha_concrete.sol|https://swcregistry.io/docs/SWC-110/|uninitialized-state|
|simple_dao.sol|https://swcregistry.io/docs/SWC-107/|reentrancy-eth|
|simple_ether_drain.sol|https://swcregistry.io/docs/SWC-105/|arbitrary-send|
|simple_suicide.sol|https://swcregistry.io/docs/SWC-106/|suicidal|
|suicide_multitx_feasible.sol|https://swcregistry.io/docs/SWC-106/|suicidal|
|suicide_multitx_infeasible.sol|https://swcregistry.io/docs/SWC-106/|suicidal|
|token-with-backdoor.sol|https://swcregistry.io/docs/SWC-110/|backdoor|
|two_mapppings.sol|https://swcregistry.io/docs/SWC-110/|uninitialized-state|
|unchecked_return_value.sol|https://swcregistry.io/docs/SWC-104/|unchecked-lowlevel|
|visibility_not_set_fixed.sol|https://swcregistry.io/docs/SWC-100/|arbitrary-send|
|visibility_not_set.sol|https://swcregistry.io/docs/SWC-100/|arbitrary-send|
|WalletLibrary.sol|https://swcregistry.io/docs/SWC-106/|reentrancy-eth, suicidal, 2 uninitialized-local|
