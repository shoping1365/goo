import axios from 'axios';

interface WeatherData {
  temp: number;
  humidity: number;
  description: string;
  icon: string;
  windSpeed: number;
  cityName: string;
}

export const useWeather = () => {
  const getWeather = async (city: string): Promise<WeatherData> => {
    try {
      const response = await axios.get(
        `https://api.openweathermap.org/data/2.5/weather?q=${city}&appid=YOUR_API_KEY&units=metric&lang=fa`
      );

      return {
        temp: Math.round(response.data.main.temp),
        humidity: response.data.main.humidity,
        description: response.data.weather[0].description,
        icon: response.data.weather[0].icon,
        windSpeed: response.data.wind.speed,
        cityName: response.data.name
      };
    } catch (error) {
      console.error('Error fetching weather data:', error);
      throw error;
    }
  };

  return {
    getWeather
  };
}; 