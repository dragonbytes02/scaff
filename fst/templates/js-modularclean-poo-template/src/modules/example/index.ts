import type {Request, Response, nextFunction} from "express"
import { MongoRepo } from "./repositories/example.mongo.repo.js";
import { CreateExampleController } from "./adapters/express/Example.controllers.js";
import { CreateExample } from "./app/createExample.js";



const examplemongoRepo = new MongoRepo()

const exampleUseCase = new CreateExample(examplemongoRepo)
const createExample = new CreateExampleController(exampleUseCase)

    
export default {
    createExample,
}