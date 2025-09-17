This RFC is being implemented. 

Activities are tracked and you are free to contribute
https://app.clickup.com/t/861mgzzwv 

Title: README.md context standardised for projects (e.g. GitLab)
Author(s): Remi Kristelijn, Marco Schriek
Version: 1.0

 

Introduction/Abstract
Stakeholders
Summary
Purpose and Outcome for personas
Motivation
Impact
Implementation
Phase 1 - Establish baseline readme.md requirements
Phase 2 - Retrofit old repositories
Alternatives
Practical example
Evaluation
1. Context
2. Prerequisites
3. Local set up
4. Validate all features including new changes
5. Release
6. More information
7. no anti-patterns
Out of scope
Anti-patterns
Examples
Related documentation
Introduction/Abstract
APS Group has more than 900 Gitlab repositories that contain projects maintained by developers. Some projects can be maintained by a singular developer, while others require entire teams. The first location that ought to give the indication of complexity and size is the repository’s readme file.

Not all repositories have a readme that gives any indication of how the project works or how to maintain it.
Additionally, the format of existing readme files varies across teams and disciplines.

Information needed to setup and test a repository is often lacking for a swift start.

Stakeholders
Application Developer:  The README.md is the first place a developer looks to work on a software project. It contains information to get started with development, deployment and relevant references/information.


https://theapsgroup.atlassian.net/wiki/pages/createpage.action?spaceKey=IMS&title=Quality%20Assurance%20Engineer: The QA team does technical analysis on projects on request. If QA has to reverse engineer a project for every evaluation, then any evaluation will take weeks and produce unreliable results. This is ineffective use of time and capacity. QA wants to focus on preventing operational risks (disruptions and low predictability) and guaranteeing business-value. Technical details are only important if they provide business value.

Summary
A good README.md sets up any persona with a staccato one page overview for:

setting up the right context;

based of the prerequisite software;

being able to install and/or run it locally;

how to validate all the code, including your bugfix, patch or new feature;

how to release/deploy it;

where to find more information.

Above 6 items are discussed below.

Purpose and Outcome for personas
Cross-team work (~10%)

A proper readme eliminates the need of other developers to make basic, maintainable changes

A readme contains a link to architecture that explains in a basic way why libraries or techniques are used in said repository

New developer (~10%)

A new developer should be able to setup the project locally and run tests.

The readme is the first place a developer will look to get themselves familiar with development or maintenance techniques for the project

Revisiting developer (~65%)

Developers often have to switch between projects/repositories. There are times that the developer doesn’t touch a project for several months or has been official off-boarded. 

When a developer comes back to work on a project, the developer may have forgotten crucial details about it. The readme must ensure that all crucial details are listed or can be discovered.

This is slightly different from the new developer as the seasoned developer looks at the information differently. With different expectations/filters, small impactful changes could be overlooked. It is also possible that a readme does describe important information, but is done poorly. This is to the same detriment as being overlooked.

Quality Assurance (~15%)

Even more than the revisiting developer, quality assurance has to switch to distinctly different context all the time to inspect different projects. The starting point needs to be to on-point, accurate and clear. Quality assurance must be able to run the project, generate and inspect reports.

Motivation
Having a sufficient README.md  provides new- or returning developers a starting point to build their own overview of what is in the repository. If the README.md is standardised this will optimise knowing where to find what in the readme files. 

We need to define what should be written in the readme as part of the standard.

Impact
Currently there are many repositories having one or more of following scenarios;

no or insufficient README.md, developers have to

ask questions to understand the project

find back which designs in Confluence-corresponds with said repository code base

dive into the code to reverse engineer (Big cognitive overhead for just getting to know a code base, error-prone)

skipping steps, like e.g.

how to get a database if you need this for a service will result in people not being able to work on the code

the prerequisites, resulting the user doesn’t have required software and therefore gets errors that don’t make sense, having to spend time fixing it

too much clutter in the README.md - It is hard to find what you need:

A readme should exist out of two types of content:

Quickstart - think of: install, test run locally

Working with the project/repository on a detailed level

some things are not important when you initially start working on the repo. Examples are: multiple ways of setting up local development; not needed, just write down the typical installation. You can always refer to the docs/ folder to write extensive documentation.

keeping the standard generated README.md from the platform, library or framework that was automatically generated when creating the initial boilerplate. This is basically the same as 'no or insufficient README.md.

Estimated time-savings:

No alignment needed with a second developer (counts as 2x timer)

Having the correct information when we needed saves context switching

No longer have to retrace what frameworks & libraries are used (several hours)

Find relationships between projects faster

Implementation
Phase 1 - Establish baseline readme.md requirements
Agree on whether README.md is a minimum quality standard for MVP-level software

Determine what is needed to adopt into scaffolding/templates

Update scaffolding/templates with Platform

Increase awareness of the scaffolding/template on read.MD

Phase 2 - Retrofit old repositories
Determine scope of retrofitting older (active) repositories

Archive super old repositories

Alternatives
No readme or incomplete readme

We need to actively chase colleague’s to get started on said project

We need to reverse engineer the code to ensure we don’t ruin the application on a detailistic level

Especially if testing doesn’t cover the entire code-base, there is no guarantee that trying to make the project to work changes the features. Think of: configuration settings, database initialisation, environment settings.

Automated scripts

We could use automated scripts to install and configure a computer for local development

We could use one line CLI scripts to run specific parts of the projects such as tests or cleanup

Developer portal

We could force projects to only use standardised pieces of infrastructure and code. Then only a reference needs to be added to the documentation explaining the standardised solutions.

If all projects use the same standards, then adding this information in the readme explicitly (as opposed to referencing it) is not required since all developers will use the standardised components at one point, where all standards would have examples, training and a subject matter expert available.

Practical example



README.md
Evaluation
The README.md should fit on one piece of paper. It should adhere to a predefined standard and it should use Gitlab-flavoured Markdown.

1. Context
Set the context: As a developer I would like to understand the system that is in the (mono)repo. There can already be things taken away from the https://theapsgroup.atlassian.net/wiki/spaces/PLAT/pages/21590605825 but there might be certain patterns applied, such as hex-pattern. Ideally this is one sentence or small diagram.

The context should also answer the question: “Why is this repository created?”. A proper response to this question is to add a business-case or link underneath this header to documentation answering this question.

2. Prerequisites
Prerequisites: As a developer I would like to understand the prerequisites that are needed in order to run the project. It should answer the questions;

“What versions of my frameworks do I need?”, e.g:

globally installed dependencies

runtime binaries

“What version of my application host do I need?”, e.g:

Java,

Node, 

Python,

.NET core

“What helper-services do I need to test my code?”, e.g:

TestContainers Docker

local RabbitMQ service

local Redis service

If you need to set up git-specific things such as login to Verdaccio, Github or Gitlab; it should refer to the documentation https://theapsgroup.atlassian.net/wiki/spaces/F/pages/10722246859.

3. Local set up
Local instance/development: As a developer I would like to get a local instance running.

If it is dependent on external systems, there should be a way to connect to them using e.g. ssh port-forwarding. The README.md should demonstrate the basic usage in a non-verbose way (staccato, to the point)

Ideally this is a very short numbered list with commands to get it up and running a.s.a.p. with no expansive explanation of what it does. A reference may be used at best. Examples are below.

Or in case of a library, or tool it should describe how to use it. e.g. put something in your index.html file, include it in your dependencies, have some additional configuration or basic examples, a terminal gif file, or something else that explains what to do.

4. Validate all features including new changes
Validating changes: As a developer I would like to change or contribute to any tests that have been set up. Whether these are syntax checks, unit tests, integration tests, regression tests or manual tests, before I propose a change I want to validate that the new or fixed feature works and everything else still works. The README.md file should clarify how to do this.

5. Release
Release the changes: As a developer I would like to understand how I release the changes to DEV, UAT/TST or production. 
• Should I set a tag? 
• should I press a button in the pipeline, after merging to main/master or should I do manual things?
The README.md should clarify these things. 

6. More information
More info: As a developer I would like to understand where I can find more information. There must be a link to confluence depicting the team responsibilities, a limited list of the features, the designs, the decisions, future plans, status of the system (legacy, in development, demo or prototype) and whom to contact. (see first point) Also if the repository has a docs/ folder, or if it uses tooling for documentation (storybook, swagger, GraphQL playground, compodoc, *.feature files or other integration tests),  it should mention it.

7. no anti-patterns
Should not have anti-patterns (described below)

Out of scope
Some recurring things can be kept out of the README.md. These recurring things are part of way of working or opinionated tools documented somewhere else. Relying on the way of working or the opinionated tools will generally introduce you how to do it. This is called work experience. We don’t document work experience, we document exceptions on that experience such as un-intuitive actions.


Prerequisites often require complex configuration that has nothing to do with the project itself. For these kind of configuration and install instructions, it transcends the purpose of the README.md or the repository. A basic reference to a confluence or vendor-specific page should be added for those instructions.
Note: Look at https://theapsgroup.atlassian.net/wiki/spaces/CM/pages/10491428882. If any of standardised tools is located there, it ensures all developer’s tools can be used in the same way regardless of platform and versioning.


The steps for setting up a git repository generally.
E.g. don’t write “clone the repository from <URL>”, assume the dev has done this already.


How to use common git commands such as (but not limited to) committing, pushing, merging and changing branches.


The list of technology this repository is dependent on
Note: This should be noted down in the repo topics.
E.g. angular-14, node-18, nestjs-8, Java-11

Anti-patterns
The repository or README.md should refrain from including the following anti-patterns;

Manual steps that could have been automated eons ago. Steps such as downloading an archive, extracting it to setup the repo should be fully automated using a script. Follow the principles of Infrastructure as Code. 


Having secrets stored in code, these should be 1) not needed and 2) stored somewhere safe if you can’t work around mocking dependencies


Extensive data or structure needed in your local databases. The minimal data should be generated using appropriate tooling, such as flyway, prisma or a self-healing startup sequence in the running application.


Not making use of the right typography e.g. hyperlinks as plain text, code as inline text so it doesn’t stand out


Making use of non-standard ways of setting up a repositories, then you have to describe exactly what you need to do.
Note: Prototypes may require this description as prototypes are not standardised and meant to proof something. A prototype could theoretically proof a non-standard way of working is better than one that is. In cases of prototypes, ALL deviating actions compared to the standards must be written down.
Else, you are not allowed to deviate.


Having to describe setting up in/on different platforms this should be abstracted away, e.g. not using backslashes in paths, non-posix way of working
Note: Keep in mind APS supports Mac, Linux, Windows, WSL. All scripts must be runnable on these platforms.


Describing engineering processes; these are managed within the team and should not be mixed with the code

Examples
To be improved

Improved example

Description

Copy the code base from gitlab.



git clone https://gitlab.apsmos.com/aps/services/spec-price-orchestrator.git
Go the project directory.



cd spec-price-orchestrator
Start the application by running the command below which requires specifying the hostname, username and password of the RabbitMQ instance you want to connect to. The values of these properties can be found in vault.



mvnw clean spring-boot:run -Dapp.rabbitmq.username=<username> -Dapp.rabbitmq.password=<password> -Dapp.rabbitmq.host=<hostname>
Note: in order to be able to connect to a remote instance of RabbitMQ, a tunnel needs first to be established. For instructions on how to do that please have a look at this page on confluence

cp src/main/resources/example.application.properties src/main/resources/application.properties - initialise the properties file and fill in the username and password

ssh -N -L 5671:tldr:5671 -D 9000 bastion-eks.mos.dev-apsmos.com open a tunnel to RabbitMQ

mvn clean spring-boot:run start the app

no need to describe to clone and change directory (anti-patterns)

having a numbered step makes clear that there are actionable steps. Having to search for this literally between the lines is error prone. Also you can say I’m stuck at step 3. 

Alternative installs can be moved to the ./docs

Related documentation
earlier propositions https://theapsgroup.atlassian.net/wiki/spaces/F/pages/1940128157 

https://theapsgroup.atlassian.net/wiki/spaces/PLAT/pages/21590605825 

What Github thinks of https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-readmes 

https://github.com/matiassingers/awesome-readme 

https://github.com/kelseyhightower/nocode 

https://www.shawnhooper.ca/2022/08/31/the-importance-of-a-readme-file-for-your-team-members-and-you/ 

https://www.freecodecamp.org/news/how-to-write-a-good-readme-file/ 