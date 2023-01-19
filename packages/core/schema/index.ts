import * as schema_bv5 from "./bv5.schema.json";
import * as schema_session from "./session.schema.json";
import Ajv from "ajv";

const ajv = new Ajv();
ajv.addSchema(schema_bv5, "bv5");
ajv.addSchema(schema_session, "session");

export { ajv };
