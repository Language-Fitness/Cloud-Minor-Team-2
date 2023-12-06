import graphene
from .Queries import Query
from .Mutations import Mutation


schema = graphene.Schema(query=Query, mutation=Mutation)