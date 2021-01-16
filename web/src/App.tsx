import React, { useEffect, useState } from 'react';

interface wasm {
  gofunc: () => number
}

const useWasm = () => {
  const [functions, setState] = useState<wasm>({
    gofunc: () => {
      return 0
    }
  })

  useEffect(() => {
    // @ts-ignore
    const go = new global.Go();
    WebAssembly.instantiateStreaming(fetch("wasm"), go.importObject).then((result) => {
      go.run(result.instance);
    }).then(() => {
      setState({
        gofunc: gofunc,
      })
    })
  }, [])

  return functions
}

export function App() {
  const wasm = useWasm()

  return (
    <div>{wasm.gofunc()}</div>
  );
}
