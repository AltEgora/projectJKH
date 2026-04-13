using Microsoft.Extensions.Options;
using GkhCalculator.Models;

namespace GkhCalculator.Services
{
    public class GkhCalculatorService
    {
        private const decimal ElectricStoveTariffFactor = 0.7m;

        private readonly GkhTariffsOptions _tariffs;

        public GkhCalculatorService(IOptions<GkhTariffsOptions> tariffsOptions)
        {
            _tariffs = tariffsOptions.Value;
        }

        public object Calculate(GkhRequest req)
        {
            var p = req.Profile;
            var residents = p?.Residents ?? req.Residents;
            if (residents < 0)
                residents = 0;

            var coldVol = ResolveColdWaterVolume(req, p, residents);
            var hotVol = ResolveHotWaterVolume(req, p, residents);

            var coldWater = coldVol * _tariffs.ColdWaterTariff;
            var hotGcal = req.HotWaterThermalGcal;
            var hotWaterSupply = hotVol * _tariffs.HotWaterVolumeTariff;
            var hotWaterHeat = hotGcal * _tariffs.HotWaterHeatTariff;
            var hotWater = hotWaterSupply + hotWaterHeat;

            var sewerVolume = coldVol + hotVol;
            var sewer = sewerVolume * _tariffs.SewerTariff;

            var electricity = ResolveElectricityCost(req, p);

            var heatVol = ResolveHeatingVolume(req, p);
            var heating = heatVol * _tariffs.HeatingTariff;

            var gasVol = ResolveGasVolume(req, p, residents);
            var gas = gasVol * _tariffs.GasTariff;

            var waste = residents * _tariffs.WasteTariffPerPerson;

            var capitalRepair = ResolveCapitalRepair(p);

            var resourceMult = CommercialMult(p);
            coldWater *= resourceMult;
            hotWaterSupply *= resourceMult;
            hotWaterHeat *= resourceMult;
            hotWater = hotWaterSupply + hotWaterHeat;
            sewer *= resourceMult;
            electricity *= resourceMult;
            heating *= resourceMult;
            gas *= resourceMult;

            var odn = ResolveOdnCharge(p, coldWater, hotWater, electricity);

            var total =
                coldWater +
                hotWater +
                sewer +
                electricity +
                heating +
                gas +
                waste +
                odn +
                capitalRepair;

            if (req.HasBenefits)
                total *= 0.7m;

            var diff = req.UserTotal - total;

            string status = Math.Abs(diff) < 1
                ? "ok"
                : diff > 0 ? "overpay" : "underpay";

            object? lineComparison = null;
            var lineMode = req.ReceiptLineMode?.Trim().ToLowerInvariant();
            if (lineMode is "cold" or "hot" or "heat")
            {
                var lineSubtotal = lineMode switch
                {
                    "cold" => coldWater + sewer,
                    "hot" => hotWaterSupply + hotWaterHeat,
                    _ => heating
                };
                var lineDiff = req.ReceiptLineAmount - lineSubtotal;
                var lineStatus = Math.Abs(lineDiff) < 1
                    ? "ok"
                    : lineDiff > 0 ? "overpay" : "underpay";
                lineComparison = new
                {
                    subtotal = lineSubtotal,
                    receiptLineAmount = req.ReceiptLineAmount,
                    difference = lineDiff,
                    status = lineStatus
                };
            }

            return new
            {
                total,
                difference = diff,
                status,
                breakdown = new
                {
                    coldWater,
                    hotWater,
                    hotWaterSupply,
                    hotWaterHeat,
                    sewer,
                    electricity,
                    heating,
                    gas,
                    waste,
                    odn,
                    capitalRepair
                },
                lineComparison
            };
        }

        private decimal CommercialMult(GkhProfile? p) =>
            p?.PropertyType == "commercial"
                ? Math.Max(1m, _tariffs.CommercialResourceMultiplier)
                : 1m;

        private decimal ResolveCapitalRepair(GkhProfile? p)
        {
            if (p == null || p.PropertyType != "residential" || p.BuildingType != "apartment")
                return 0m;
            if (p.Area <= 0m)
                return 0m;
            return p.Area * _tariffs.CapitalRepairTariffPerM2;
        }

        private decimal ResolveColdWaterVolume(GkhRequest req, GkhProfile? p, int residents)
        {
            if (p != null && !p.WaterMeter && residents > 0)
                return residents * _tariffs.ColdWaterNormM3PerPerson;
            return req.ColdWaterVolume;
        }

        private decimal ResolveHotWaterVolume(GkhRequest req, GkhProfile? p, int residents)
        {
            if (p != null && !p.WaterMeter && residents > 0)
                return residents * _tariffs.HotWaterNormM3PerPerson;
            return req.HotWaterVolume;
        }

        private decimal ResolveElectricityCost(GkhRequest req, GkhProfile? p)
        {
            var dayTariff = _tariffs.ElectricityTariff;
            var nightTariff = _tariffs.ElectricityNightTariff;
            if (p?.StoveType == "electric")
            {
                dayTariff *= ElectricStoveTariffFactor;
                nightTariff *= ElectricStoveTariffFactor;
            }

            var day = req.ElectricityDayVolume;
            var night = req.ElectricityNightVolume;
            var sumDayNight = day + night;

            if (p?.ElectricityType == "dual")
                return day * dayTariff + night * nightTariff;

            var kwh = sumDayNight > 0 ? sumDayNight : req.ElectricityVolume;
            return kwh * dayTariff;
        }

        private decimal ResolveHeatingVolume(GkhRequest req, GkhProfile? p)
        {
            if (p?.HeatingType == "none")
                return 0m;

            if (p != null && !p.HeatingMeter && p.Area > 0 && p.HeatingType != "none")
                return p.Area * _tariffs.HeatingNormGcalPerM2;

            return req.HeatingVolume;
        }

        private decimal ResolveGasVolume(GkhRequest req, GkhProfile? p, int residents)
        {
            if (p?.GasType == "none")
                return 0m;
            if (p?.GasType == "norm" && residents > 0)
                return residents * _tariffs.GasNormM3PerPerson;
            return req.GasVolume;
        }

        private decimal ResolveOdnCharge(GkhProfile? p, decimal coldWater, decimal hotWater, decimal electricity)
        {
            if (p == null || !p.Odn || p.BuildingType != "apartment")
                return 0m;
            var baseSum = coldWater + hotWater + electricity;
            return baseSum * _tariffs.OdnFraction;
        }
    }
}
