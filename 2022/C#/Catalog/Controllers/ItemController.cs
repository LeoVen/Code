using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using AutoMapper;
using Catalog.Dtos;
using Catalog.Entitities;
using Catalog.Repositories;
using Microsoft.AspNetCore.Mvc;

namespace Catalog.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class ItemController : ControllerBase
    {
        private readonly IItemsRepository repository;
        private readonly IMapper mapper;

        public ItemController(IMapper mapper, IItemsRepository repository)
        {
            this.repository = repository;
            this.mapper = mapper;
        }

        [HttpGet]
        [Route(nameof(GetItems))]
        public async Task<ActionResult<IEnumerable<ItemDto>>> GetItems()
        {
            var items = await repository.GetItemsAsync();

            return Ok(mapper.Map<IEnumerable<Item>, IEnumerable<ItemDto>>(items));
        }

        [HttpGet("[action]/{id}")]
        [Route(nameof(GetItem))]
        public async Task<ActionResult<ItemDto>> GetItem(Guid id)
        {
            var item = await repository.GetItemAsync(id);

            if (item is null)
                return NotFound();

            return Ok(mapper.Map<Item, ItemDto>(item));
        }

        [HttpPost]
        [Route(nameof(CreateItem))]
        public async Task<ActionResult<ItemDto>> CreateItem(CreateItemDto itemDto)
        {
            // CreateItemDto -> Item -> ItemDto
            var toInsert = mapper.Map<CreateItemDto, Item>(itemDto);

            await repository.CreateItemAsync(toInsert);

            var result = mapper.Map<Item, ItemDto>(toInsert);

            return CreatedAtAction(nameof(GetItem), new { Id = result.Id }, result);
        }

        [HttpPut]
        [Route(nameof(UpdateItem))]
        public async Task<ActionResult> UpdateItem([FromBody] UpdateItemDto updateItem)
        {
            var existingItem = repository.GetItemAsync(updateItem.Id);

            if (existingItem is null)
                return NotFound();

            var newItem = mapper.Map<UpdateItemDto, Item>(updateItem);

            await repository.UpdateItemAsync(newItem);

            return NoContent();
        }

        [HttpDelete("[action]/{id}")]
        [Route(nameof(DeleteItem))]
        public async Task<ActionResult> DeleteItem(Guid id)
        {
            var existingItem = await repository.GetItemAsync(id);

            if (existingItem is null)
                return NotFound();

            await repository.DeleteItemAsync(existingItem.Id);

            return NoContent();
        }
    }
}
