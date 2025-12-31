import {Request,Response,NextFunction} from 'express'
import { CreateExampleUseCase } from '../../../internal/useCases/example/Create.js'
import { ExamplePostgresRepository } from '../../../internal/repositories/example/postgres.js'
import { CreateExampleController } from '../../../adapters/http/example/Create.js'


const exampleRepositoryAssembly = new ExamplePostgresRepository()

const createExampleUseCase = new CreateExampleUseCase(exampleRepositoryAssembly)

const createExample = new CreateExampleController(createExampleUseCase)

export default { 
    createExample, 
}