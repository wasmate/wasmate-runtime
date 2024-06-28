use super::utils::{deploy_contract};
use wasmate::WasmCidManager::{IWasmCidManagerDispatcher, IWasmCidManagerDispatcherTrait,cidMetaData};

//
 // @function test_add_wasm_cid
 // @description This function tests the add_wasm_cid method of the IWasmCidManagerDispatcher.
 // It checks if the method successfully adds a WASM CID to the contract.
 // @param {number} initial_counter - The initial counter value used to deploy the contract.
 // @param {string} contract_address - The address of the deployed contract.
 // @param {string} faas_func_name - The name of the FAAS function to which the WASM CID will be added.
 // @param {object} wasm_cid_meta_data - The metadata of the WASM CID to be added.
 // @param {boolean} expected_result - The expected result of the add_wasm_cid method.
 // @returns {boolean} - Returns true if the add_wasm_cid method succeeds, false otherwise.
//
#[test]
    fn test_add_wasm_cid() {
        let initial_counter = 0;
        let contract_address = deploy_contract(initial_counter);
        // Define the FAAS function name, WASM CID metadata, and expected result
        let faas_func_name = 'sayhello';
        let wasm_cid_meta_data = cidMetaData {
            cid_address_part1: 'oldpart1',
            cid_address_part2: 'oldpart2',
        };
        let expected_result = true;

        let dispatcher = IWasmCidManagerDispatcher { contract_address };
        let isSuccess=dispatcher.add_wasm_cid(faas_func_name,wasm_cid_meta_data);
        // Call the add_wasm_cid function and assert the result
        assert!(isSuccess == expected_result, "add_wasm_cid fail.");
}


//
 // @function test_get_wasm_cid
 // @description This function tests the get_wasm_cid method of the IWasmCidManagerDispatcher.
 // It checks if the method successfully retrieves the WASM CID metadata for a given FAAS function.
 // @param {number} initial_counter - The initial counter value used to deploy the contract.
 // @param {string} contract_address - The address of the deployed contract.
 // @param {string} faas_func_name - The name of the FAAS function for which the WASM CID metadata is to be retrieved.
 // @returns {boolean} - Returns true if the get_wasm_cid method successfully retrieves the WASM CID metadata, false otherwise.
//
#[test]
    fn test_get_wasm_cid() {
    let initial_counter = 0;
    let contract_address = deploy_contract(initial_counter);
    // Define the FAAS function name, WASM CID metadata, and expected result
    let faas_func_name = 'sayhello';
    let wasm_cid_meta_data = cidMetaData {
        cid_address_part1: 'oldpart1',
        cid_address_part2: 'oldpart2',
    };

    let dispatcher = IWasmCidManagerDispatcher { contract_address };
    let isSuccess=dispatcher.add_wasm_cid(faas_func_name,wasm_cid_meta_data);
    
    assert!(isSuccess == true , "add_wasm_cid fail.");

    let getwasmcid= dispatcher.get_wasm_cid(faas_func_name);

    assert!(getwasmcid.cid_address_part1 == wasm_cid_meta_data.cid_address_part1, "get_wasm_cid part1 fail.");
    assert!(getwasmcid.cid_address_part2 == wasm_cid_meta_data.cid_address_part2, "get_wasm_cid part2 fail.");
}

//
 // @function test_update_wasm_cid
 // @description This function tests the update_wasm_cid method of the IWasmCidManagerDispatcher.
 // It checks if the method successfully updates the WASM CID metadata for a given FAAS function.
 // @param {number} initial_counter - The initial counter value used to deploy the contract.
 // @param {string} contract_address - The address of the deployed contract.
 // @param {string} faas_func_name - The name of the FAAS function for which the WASM CID metadata is to be updated.
 // @param {object} wasm_cid_meta_data - The new metadata of the WASM CID to be updated.
 // @param {boolean} expected_result - The expected result of the update_wasm_cid method.
 // @returns {boolean} - Returns true if the update_wasm_cid method successfully updates the WASM CID metadata, false otherwise.
//
#[test]
    fn test_update_wasm_cid() {
    let initial_counter = 0;
    let contract_address = deploy_contract(initial_counter);
    // Define the FAAS function name, WASM CID metadata, and expected result
    let faas_func_name = 'sayhello';
    let wasm_cid_meta_data = cidMetaData {
        cid_address_part1: 'oldpart1',
        cid_address_part2: 'oldpart2',
    };

    let dispatcher = IWasmCidManagerDispatcher { contract_address };
    let isSuccess=dispatcher.add_wasm_cid(faas_func_name,wasm_cid_meta_data);
    
    assert!(isSuccess == true , "add_wasm_cid fail.");


    let wasm_cid_meta_data_new = cidMetaData {
        cid_address_part1: 'newpart1',
        cid_address_part2: 'newpart2',
    };

    
    let isSuccess= dispatcher.update_wasm_cid(faas_func_name,wasm_cid_meta_data_new);
    assert!(isSuccess == true, "update_wasm_cid fail.");

    let getwasmcid= dispatcher.get_wasm_cid(faas_func_name);

    assert!(getwasmcid.cid_address_part1 == wasm_cid_meta_data_new.cid_address_part1, "update_wasm_cid part1 fail.");
    assert!(getwasmcid.cid_address_part2 == wasm_cid_meta_data_new.cid_address_part2, "update_wasm_cid part2 fail.");
}

//
// @function check_get_wasmcid_total_count
// @description This function tests the get_wasmcid_total_count method of the IWasmCidManagerDispatcher.
// It checks if the total count of WASM CIDs stored in the contract matches the initial counter.
// @param {number} initial_counter - The initial counter value used to deploy the contract.
// @param {string} contract_address - The address of the deployed contract.
// @returns {boolean} - Returns true if the stored counter matches the initial counter, false otherwise.
//
#[test]
fn check_get_wasmcid_total_count() {
    let initial_counter = 0;
    let contract_address = deploy_contract(initial_counter);
    let dispatcher = IWasmCidManagerDispatcher { contract_address };
    let stored_counter = dispatcher.get_wasmcid_total_count();
    assert!(stored_counter == initial_counter, "get_wasmcid_total_count fail.");
}