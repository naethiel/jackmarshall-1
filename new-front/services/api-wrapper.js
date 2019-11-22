export function callApi(url, payload, method = "GET") {
  return fetch(url, {
    credentials: "omit",
    method: method,
    mode: "cors",
    body: JSON.stringify(payload) // body data type must match "Content-Type" header
  });
}
