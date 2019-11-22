import { callApi } from "./api-wrapper";
import config from "../config";

export function login(username, password) {
  return callApi(`${config.auth}/login`, { login: username, password }, "POST");
}
