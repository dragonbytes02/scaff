import type { Example } from "../domain/entities.js";
import type { DbRepository } from "../ports/db.repository.js";


export class MongoRepo  implements DbRepository{

    create(example: Example): Example {
        throw new Error("Method not implemented.");
    }
    update(example: Example): Example {
        throw new Error("Method not implemented.");
    }

}