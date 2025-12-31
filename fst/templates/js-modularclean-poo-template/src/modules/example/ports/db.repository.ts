import type { Example } from "../domain/entities.js"

export interface DbRepository  {

    create(example:Example):Example
    update(example:Example):Example
}