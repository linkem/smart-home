using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;
using System;
using System.Collections.Generic;

namespace Weather.Domain
{
    public class WeatherEntity
    {
        [BsonId]
        [BsonRepresentation(BsonType.ObjectId)]
        public string ID { get; set; }

        public float Temp { get; set; }
        public float Humidity { get; set; }

        public IEnumerable<WeatherConditionEntity> Conditions { get; set; }

        public string CityName { get; set; }
        public int CityId { get; set; }

        public float Clouds { get; set; }
        [BsonDateTimeOptions(Kind = DateTimeKind.Utc)]
        public DateTime DateTimeUtc { get; set; }
        public DateTime SunsetDateTImeUtc { get; set; }
        public DateTime SunriseDateTImeUtc { get; set; }
    }

    public class WeatherConditionEntity
    {
        public int Id { get; set; }
        public string Main { get; set; }
        public string Description { get; set; }
        public string Icon { get; set; }
    }
}
