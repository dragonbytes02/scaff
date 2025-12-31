import type { Example } from "../domain/entities.js";

export interface CreateExampleUseCase {
    execute (data: Example):Example
}