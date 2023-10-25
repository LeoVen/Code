import { PrismaClient, Provider, ProviderType, Relationship, Solution } from '@prisma/client'

interface CreateProviderDto {
    name: string,
    type: ProviderType,
}

interface CreateSolutionDto {
    name: string,
    solutionID: string
}

class Repository {
    constructor() { }

    async createProvider(provider: CreateProviderDto): Promise<Provider> {
        return await prisma.provider.create({
            data: {
                name: provider.name,
                type: provider.type,
            }
        })
    }

    async createSolution(solution: CreateSolutionDto): Promise<Solution> {
        return await prisma.solution.create({
            data: {
                name: solution.name,
                solutionID: solution.solutionID,
            }
        })
    }

    async createRelationship(rel: Relationship): Promise<Relationship> {
        return await prisma.relationship.create({
            data: {
                providerId: rel.providerId,
                solutionId: rel.solutionId,
            }
        })
    }

    async updateProvider(provider: Provider): Promise<Provider> {

        return await prisma.provider.update({
            where: {
                id: provider.id,
            },
            data: {
                name: provider.name,
                type: provider.type,
            }
        })
    }

    async updateSolution(solution: Solution): Promise<Solution> {

        return await prisma.solution.update({
            where: {
                id: solution.id,
            },
            data: {
                name: solution.name,
                solutionID: solution.solutionID,
            }
        })
    }

    async listProviders(): Promise<Provider[]> {
        return await prisma.provider.findMany({
            select: {
                id: true,
                name: true,
                type: true,
                Relationship: {
                    include: {
                        solution: true,
                    }
                },
            }
        })
    }

    async listSolutions(): Promise<Solution[]> {
        return await prisma.solution.findMany({
            select: {
                id: true,
                name: true,
                solutionID: true,
                Relationship: {
                    include: {
                        provider: true,
                    }
                },
            }
        })
    }

    async listSolutionsByProvider(provider: Provider): Promise<Solution[]> {
        return prisma.solution.findMany({
            where: {
                Relationship: {
                    some: {
                        providerId: provider.id,
                    }
                }
            }
        })
    }

    async providerByName(name: string): Promise<Provider | null> {
        return await prisma.provider.findFirst({
            where: {
                name: name,
            },
            select: {
                id: true,
                name: true,
                type: true,
            }
        })
    }

    async solutionByName(name: string): Promise<Solution | null> {
        return await prisma.solution.findFirst({
            where: {
                name: name,
            },
            select: {
                id: true,
                name: true,
                solutionID: true,
            }
        })
    }

    async eraseDatabase(): Promise<void> {
        await prisma.provider.deleteMany()
        await prisma.solution.deleteMany()
    }

    async deleteProvider(provider: Provider): Promise<void> {
        await prisma.provider.delete({
            where: {
                id: provider.id,
            }
        })
    }
}

const prisma = new PrismaClient({ log: ['query'] });

async function main() {

    let repo = new Repository()

    // 0
    // await repo.eraseDatabase();

    // 1
    // let provider = await repo.createProvider({ name: 'Provider1', type: 'EXTERNAL' })
    // let solution = await repo.createSolution({ name: 'Solution1', solutionID: 'd679d3cd-a8c7-43fb-82e4-89260db5ae6b' })

    // 2
    // let provider = await repo.providerByName('Provider1')
    // let solution = await repo.solutionByName('Solution1')
    // if (provider && solution) {
    //     await repo.createRelationship({ providerId: provider.id, solutionId: solution.id })
    // }

    // 3
    let provider = await repo.providerByName('Provider1')
    if (provider) {
        let solutions = await repo.listSolutionsByProvider(provider)
        console.log(solutions)
    }

    // console.log('--------------------------------------')
    // console.log(await repo.listProviders())
    // console.log('--------------------------------------')
    // console.log(await repo.listSolutions())
}

main();
