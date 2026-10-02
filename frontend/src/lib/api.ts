// TODO: Add browser-side API functions after defining docs/api.yaml.
// Keep request construction and response handling here, separate from the UI.
// Define response types from your actual contract; handle failed responses.
// No requests or response shapes are implemented yet.
export async function checkAPI(): Promise<string>{
  const response = await fetch('http://localhost:8080/healthz');

  if(!response.ok){
    throw new Error("Error: " + response.status);
  }

  const data = await response.json();
  return data.message;
  
}
