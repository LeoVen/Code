import { PrismaClient, ProviderType, Provider, Solution } from '@prisma/client'

interface CreateProviderDto {
    name: string,
    type: ProviderType,
    solutionIDs: string[]
}

interface CreateSolutionDto {
    name: string,
    solutionID: string
    providerIDs: string[]
}

class Repository {
    constructor() { }

    async createProvider(provider: CreateProviderDto): Promise<Provider> {
        return await prisma.provider.create({
            data: {
                name: provider.name,
                type: provider.type,
                solutionIDs: provider.solutionIDs,
            }
        })
    }

    async createSolution(solution: CreateSolutionDto): Promise<Solution> {
        return await prisma.solution.create({
            data: {
                name: solution.name,
                solutionID: solution.solutionID,
                providerIDs: solution.providerIDs,
            }
        })
    }

    async assignRelationship(provider: Provider, solution: Solution): Promise<void> {

        let update = false
        if (!provider.solutionIDs.includes(solution.id)) {
            provider.solutionIDs.push(solution.id)
            update = true
        }
        if (!solution.providerIDs.includes(provider.id)) {
            solution.providerIDs.push(provider.id)
            update = true
        }

        if (update) {
            await this.updateProvider(provider)
            await this.updateSolution(solution)
        }
    }

    async updateProvider(provider: Provider): Promise<Provider> {
        return await prisma.provider.update({
            where: {
                id: provider.id,
            },
            data: {
                name: provider.name,
                type: provider.type,
                solutionIDs: provider.solutionIDs,
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
                providerIDs: solution.providerIDs,
            }
        })
    }

    async listProviders(): Promise<Provider[]> {
        return await prisma.provider.findMany({
            select: {
                id: true,
                name: true,
                type: true,
                solutions: true,
                solutionIDs: true,
            }
        })
    }

    async listSolutions(): Promise<Solution[]> {
        return await prisma.solution.findMany({
            select: {
                id: true,
                name: true,
                solutionID: true,
                providers: true,
                providerIDs: true,
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
                solutions: true,
                solutionIDs: true,
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
                providers: true,
                providerIDs: true,
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

const prisma = new PrismaClient();

async function main() {

    let repo = new Repository()

    await repo.createProvider({
        name: 'Prov1', type: 'EXTERNAL', solutionIDs: [],
    })
    await repo.createSolution({
        name: 'Solar', solutionID: 'cffb9354-d993-42e7-aa18-0f58d3726ce5', providerIDs: [],
    })

    console.log('--------------------------------------')
    console.log(await repo.listProviders())
    console.log('--------------------------------------')
    console.log(await repo.listSolutions())
}

main();
