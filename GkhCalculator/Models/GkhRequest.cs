namespace GkhCalculator.Models
{
    public class GkhRequest
    {
        public decimal ColdWaterVolume { get; set; }
        public decimal SewerVolume { get; set; }
        public decimal ElectricityVolume { get; set; }
        public decimal HeatingVolume { get; set; }

        public bool HasBenefits { get; set; }

        public decimal UserTotal { get; set; }
    }
}