import type { Request, Response , NextFunction } from "express";
import type { CreateExampleUseCase } from "../../ports/example.services.js";



export class CreateExampleController {

    constructor(private example: CreateExampleUseCase){}

    execute(req:Request, res:Response, nextFunction:NextFunction){

        const result = this.example.execute(req.body)

        res.status(201).json(result)
    }
}
