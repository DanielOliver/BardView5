import axios, {AxiosResponse} from "axios";
import {SelfServiceRegistrationFlow} from "@ory/kratos-client";

async function startSelfServiceRegister(): Promise<SelfServiceRegistrationFlow> {
  const response: AxiosResponse<SelfServiceRegistrationFlow> = await axios.get('/self-service/registration/api');
  return response.data
}

export {
  startSelfServiceRegister
}
