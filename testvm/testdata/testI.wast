(module
  (import "debug" "print32" (func $log (param i32)))
  (func $print32 (param $i i32)
    get_local $i
    call $log)
  (export "print32" (func $print32)))