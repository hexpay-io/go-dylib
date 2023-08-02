use std::ffi::{c_char, CString};

#[no_mangle]
pub extern "C" fn example() -> *mut c_char {
    CString::new("Hello from Rust!")
        .unwrap()
        .into_raw()
}

#[no_mangle]
pub unsafe extern "C" fn free_c_char(ptr: *mut c_char) {
    if ptr.is_null() {
        return;
    }
    let _ = CString::from_raw(ptr);
}