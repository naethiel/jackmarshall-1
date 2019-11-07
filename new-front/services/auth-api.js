import { callApi } from "./api-wrapper";
import config from "../config";

export function login(username, password) {
  return callApi(`${config.auth}/login`, { username, password }, "POST");
}
