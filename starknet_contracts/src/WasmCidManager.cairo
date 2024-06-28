#[starknet::interface]
// The name of the interface  to be used
trait IWasmCidManager<TContractState> {
    /////
     // Adds a new WASM CID resource along with related metadata.
     //
     // @param self The contract state instance.
     // @param faas_func_name The name of the FAAS (Function as a Service) resource.
     // @param wasm_cid_meta_data The metadata associated with the WASM CID.
     // @returns True if the operation was successful, false otherwise.
     ///
    fn add_wasm_cid(ref self: TContractState,faas_func_name:felt252, wasm_cid_meta_data:cidMetaData) -> bool;

    /////
     // Retrieves metadata associated with a given WASM CID.
     //
     // @param self The contract state instance.
     // @param faas_func_name The name of the FAAS (Function as a Service) resource.
     // @returns The metadata associated with the specified WASM CID.
     ///
    fn get_wasm_cid(self: @TContractState,faas_func_name:felt252) -> cidMetaData;

    /////
     // Updates the metadata for a specified WASM CID.
     //
     // @param self The contract state instance.
     // @param faas_func_name The name of the FAAS (Function as a Service) resource.
     // @param wasm_cid_meta_data The updated metadata associated with the WASM CID.
     // @returns True if the operation was successful, false otherwise.
     ///
    fn update_wasm_cid(ref self: TContractState,faas_func_name:felt252, wasm_cid_meta_data:cidMetaData) -> bool;

    /////
     // get_wasmcid_totoal_count is the number of wasmcid to be used.
     //
     // @param self The contract state instance.
     // @returns The total count of WASM CIDs.
     ///
    fn get_wasmcid_total_count(self: @TContractState) -> u128;
}

#[derive(Drop, Serde, Copy, starknet::Store)]
/// A struct representing the metadata for a WASM CID.
/// It contains two fields: `cid_address_part1` and `cid_address_part2`, which together form the WASM CID address.
pub struct cidMetaData {
    /// The first part of the WASM CID address.
    pub cid_address_part1: felt252,
    /// The second part of the WASM CID address.
    pub cid_address_part2: felt252,
}

#[starknet::contract]
pub mod WasmCidManager {
    use starknet::ContractAddress;
    use super::cidMetaData;

    /// The storage struct for the WasmCidManager contract. It contains two fields:
    /// `wasmcidTotalCount` which is the total count of WASM CIDs, and
    /// `wasmCidData` which is a LegacyMap that stores the metadata for each WASM CID.
    #[storage]
    struct Storage {
        wasmcidTotalCount: u128,
        wasmCidData: LegacyMap::<felt252, cidMetaData>,
    }

    #[constructor]
    fn constructor(ref self: ContractState, init_value: u128) {
        self.wasmcidTotalCount.write(0);
    }

    #[abi(embed_v0)]
    impl WasmCidManager of super::IWasmCidManager<ContractState> {
        /// Adds a new WASM CID resource along with related metadata.
        ///
        /// @param self The contract state instance.
        /// @param faas_func_name The name of the FAAS (Function as a Service) resource.
        /// @param wasm_cid_meta_data The metadata associated with the WASM CID.
        /// @returns True if the operation was successful, false otherwise.
        ///
        fn add_wasm_cid(ref self: ContractState,faas_func_name:felt252, wasm_cid_meta_data:cidMetaData) -> bool {

            self.wasmCidData.write(faas_func_name, wasm_cid_meta_data);

            let totalCounter = self.wasmcidTotalCount.read() + 1;
            self.wasmcidTotalCount.write(totalCounter);
            return true;
        }

        /// Retrieves metadata associated with a given WASM CID.
        ///
        /// @param self The contract state instance.
        /// @param faas_func_name The name of the FAAS (Function as a Service) resource.
        /// @returns The metadata associated with the specified WASM CID.
        ///
        fn get_wasm_cid(self: @ContractState,faas_func_name:felt252) -> cidMetaData {
            return self.wasmCidData.read(faas_func_name);
        }       

        /// Updates the metadata for a specified WASM CID.
        ///
        /// @param self The contract state instance.
        /// @param faas_func_name The name of the FAAS (Function as a Service) resource.
        /// @param wasm_cid_meta_data The updated metadata associated with the WASM CID.
        /// @returns True if the operation was successful, false otherwise.
        ///
        fn update_wasm_cid(ref self: ContractState,faas_func_name:felt252, wasm_cid_meta_data:cidMetaData) -> bool {
           
           self.wasmCidData.write(faas_func_name, wasm_cid_meta_data);
           
            return true;
        }

        /// get_wasmcid_total_count is the number of wasmcid to be used.
        ///
        /// @param self The contract state instance.
        /// @returns The total count of WASM CIDs.
        ///
        fn get_wasmcid_total_count(self: @ContractState) -> u128 {
            return self.wasmcidTotalCount.read();
        }

    }
}
