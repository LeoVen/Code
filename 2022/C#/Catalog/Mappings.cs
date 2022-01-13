using System;
using AutoMapper;
using Catalog.Dtos;
using Catalog.Entitities;

namespace Catalog
{
    public class Mappings : Profile
    {
        public Mappings()
        {
            CreateMap<Item, ItemDto>()
                .ReverseMap();

            CreateMap<CreateItemDto, Item>()
                .ForMember(item => item.Id, map => map.MapFrom(_ => Guid.NewGuid()))
                .ForMember(item => item.CreatedDate, map => map.MapFrom(_ => DateTimeOffset.UtcNow));

            CreateMap<UpdateItemDto, Item>()
                .ForMember(item => item.CreatedDate, map => map.MapFrom(_ => DateTimeOffset.UtcNow));
        }
    }
}
