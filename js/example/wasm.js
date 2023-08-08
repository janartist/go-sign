'use strict';

const WASM_URL = 'main.wasm';

var wasm;


function init() {

  const go = new Go();
  if ('instantiateStreaming' in WebAssembly) {
    WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function (obj) {
      wasm = obj.instance;
      go.run(wasm);

      // 调用函数
      var res = testAlert();
      console.log(res)

      // 示例数据对象
      const data = {
        value: 42,
        val:"wo de我",
      };

      // 示例配置对象
      const config = {
        t: 12345,
        nonce: "wefw23swdwef",
        secret: 'mysecret',
      };
      console.log("Calling signEncode with data:", data);
      console.log("Calling signEncode with config:", config);
      var res = signEncode.call(this, data, config);
      console.log(res)

    })
  } else {
    fetch(WASM_URL).then(resp =>
        resp.arrayBuffer()
    ).then(bytes =>
        WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
          wasm = obj.instance;
          go.run(wasm);
        })
    )
  }
}

init();