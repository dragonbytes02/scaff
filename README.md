 # pragmatic architectures
 
 Clean Architecture, DDD, Hexagonal Architecture, MVC, Modular, Onion... Enough already!

Software architectures are fine, but often, for many projects, they become an exercise in excessive reengineering, where you'll waste a lot of time implementing code you don't need. These complex architectures are designed to support large development teams working simultaneously, which is why they are divided into layers and have so many formalities.

When you assimilate the fundamental principles of all these architectures, you realize that applying them partially or adapting them is more than enough for many projects.

# short:

Applying Clean Architecture to a small REST API is a complete waste of time. However, not applying its principles to your long-term project (a startup, for example) becomes technical debt that, sooner or later, you will have to pay dearly for.

That is exactly what Scaff seeks to do: apply the principles of these architectures while simplifying many aspects.

Scaff offers you different project structures, minimalist configurations, and best practices. All of this is designed to help you start developing your code as quickly as possible. We are not a framework; developers are given complete freedom to choose the technologies they want. If you look at Scaff's templates, they are designed so that the developer can decide how much complexity they really need in their project, simplifying the way the code is written.


> Do you need a class for each use case?
> Or would you prefer a single class that handles all your use cases?

**Go for it!**

> Do you think you'll never change your database? 
> Why implement each Repository if you're not going to use it?

 **Go for it!**

> Are you sure you'll never change your server? 
>Why inject use cases into the controllers? Make it depend directly on them.

 **Go for it!**

**You have the freedom to do whatever you want!**

But first... Learn the principles behind all these architectures. Only then will you understand exactly what you gain and what you lose with each simplification.


# We invite you to participate in this project!

Review the different templates and create your pull request with an improvement or a new structure.

To ensure that contributions maintain Scaff's philosophy, your pull request will be evaluated under the following key criteria:

-   **Balance between Simplicity and Principles:** The proposal must maintain pragmatic simplicity, but without sacrificing fundamental principles such as independence (decoupling), testability, and maintainability.
    
-   **Flexibility and Choice:** The structure should give the developer the possibility to opt for simpler solutions or evolve towards more complex implementations according to the actual needs of the project.
    
-   **Generic Configuration:** Initial configurations should be minimalist and designed to be applicable to any type of project without imposing specific technologies.



# Typescript Templates

>These are not literal implementations of clean architectures, but rather simplifications. Therefore, many elements have been intentionally sacrificed.

**Explore the different templates and read their descriptions to understand!**

[Clean Arquitecture](https://github.com/dragonbytes02/scaff/edit/main/README.md)
