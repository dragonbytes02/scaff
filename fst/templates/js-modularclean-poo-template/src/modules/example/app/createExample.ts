import type { Example } from "../domain/entities.js";
import type { DbRepository } from "../ports/db.repository.js";
import type { CreateExampleUseCase } from "../ports/example.services.js";


export class CreateExample  implements CreateExampleUseCase {

    constructor(private repo :DbRepository){}
    execute(example: Example): Example {
        throw new Error("Method not implemented.");
    }
    
}