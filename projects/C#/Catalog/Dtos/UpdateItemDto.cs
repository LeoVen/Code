using System;
using System.ComponentModel.DataAnnotations;

namespace Catalog.Dtos
{
    public record UpdateItemDto
    {
        [Required]
        public Guid Id { get; init; }
        [Required]
        public string Name { get; init; }
        [Required]
        [Range(1, int.MaxValue)]
        public decimal Price { get; init; }
    }
}
