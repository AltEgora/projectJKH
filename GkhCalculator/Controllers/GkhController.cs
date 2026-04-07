using Microsoft.AspNetCore.Mvc;
using GkhCalculator.Services;
using GkhCalculator.Models;

namespace GkhCalculator.Controllers
{
    [ApiController]
    [Route("api/gkh")]
    public class GkhController : ControllerBase
    {
        private readonly GkhCalculatorService _service;

        public GkhController(GkhCalculatorService service)
        {
            _service = service;
        }

        [HttpPost("calculate")]
        public IActionResult Calculate([FromBody] GkhRequest request)
        {
            var result = _service.Calculate(request);
            return Ok(result);
        }
    }
}