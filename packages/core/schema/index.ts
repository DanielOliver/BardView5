import * as schema_bv5 from "./bv5.schema.json";
import * as schema_session from "./session.schema.json";
import Ajv from "ajv";
import { Bv5, Bv5Session } from "../open/classes";

const ajv = new Ajv();

const validate_bv5 = ajv.compile<Bv5>(schema_bv5);
const validate_session = ajv.compile<Bv5Session>(schema_session);

export { validate_bv5, validate_session };
