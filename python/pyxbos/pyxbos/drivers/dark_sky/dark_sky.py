from pyxbos import *
import os,sys
import json
import requests
import yaml
import argparse




class DarkSkyPredictionDriver(Driver):
    def setup(self, cfg):
        self.baseurl = cfg['darksky']['url']
        self.apikey = cfg['darksky']['apikey']
        self.coords = cfg['darksky']['coordinates']
        self.url = self.baseurl + self.apikey + '/' + self.coords

    def read(self, requestid=None):
        response = requests.get(self.url)
        json_data = json.loads(response.text)
        if 'hourly' not in json_data: return

        hourly = json_data['hourly']

        predictions = []

        for hour in hourly.get('data',[]):
            timestamp = int(hour.get('time') * 1e9) # nanoseconds
            temperature = hour.get('apparentTemperature', None)
            precipIntensity = hour.get('precipIntensity', None)
            precipProbability = hour.get('precipProbability', None)
            humidity = hour.get('humidity', None)
            if humidity is not None:
                humidity *= 100 # change from decimal to percent

            predictions.append(iot_pb2.WeatherStationPrediction.Prediction(
                prediction_time=timestamp,
                prediction=iot_pb2.WeatherStation(
                    temperature=types.Double(value=temperature),
                    precip_intensity=types.Double(value=precipIntensity),
                    humidity=types.Double(value=humidity),
                )
            ))

        msg = xbos_pb2.XBOS(
            XBOSIoTDeviceState = iot_pb2.XBOSIoTDeviceState(
                time = int(time.time()*1e9),
                weather_station_prediction = iot_pb2.WeatherStationPrediction(
                    predictions=predictions
                )
            )
        )
        self.report(self.coords+'/prediction', msg)

class DarkSkyDriver(Driver):
    def setup(self, cfg):
        self.baseurl = cfg['darksky']['url']
        self.apikey = cfg['darksky']['apikey']
        self.coords = cfg['darksky']['coordinates']
        self.url = self.baseurl + self.apikey + '/' + self.coords

    def read(self, requestid=None):
        response = requests.get(self.url)
        json_data = json.loads(response.text)
        if 'currently' not in json_data: return

        logging.info("currently {0}".format(json_data['currently']))
        nearestStormDistance =  json_data['currently'].get('nearestStormDistance',None)
        nearestStormBearing =   json_data['currently'].get('nearestStormBearing',None)
        precipIntensity =       json_data['currently'].get('precipIntensity',None)
        apparentTemperature =   json_data['currently'].get('apparentTemperature',None)
        humidity =              json_data['currently'].get('humidity',None)
        if humidity is not None:
            humidity *= 100 # change from decimal to percent

        msg = xbos_pb2.XBOS(
            XBOSIoTDeviceState = iot_pb2.XBOSIoTDeviceState(
                time = int(time.time()*1e9),
                weather_station = iot_pb2.WeatherStation(
                    nearest_storm_distance  =   types.Double(value=nearestStormDistance),
                    nearest_storm_bearing   =   types.Int32(value=nearestStormBearing),
                    precip_intensity        =   types.Double(value=precipIntensity),
                    temperature             =   types.Double(value=apparentTemperature),
                    humidity                =   types.Double(value=humidity),
                )
            )
        )
        self.report('blr', msg)


if __name__ == '__main__':

    with open('dark_sky.yaml') as f:
        # use safe_load instead load for security reasons
        driverConfig = yaml.safe_load(f)

    api_key = driverConfig['dark_sky']['api_key']
    namespace = driverConfig['wavemq']['namespace']

    cfg = {
        'darksky': {
            'apikey': api_key,
            'url': 'https://api.darksky.net/forecast/',
            'coordinates': '40.5301,-124.0000' # Should be near BLR
        },
        'wavemq': 'localhost:4516',
        'namespace': namespace,
        'base_resource': 'dark_sky',
        'entity': 'test_dark_sky.ent',
        'id': 'pyxbos-driver-darksky-1',
        #'rate': 1800, # half hour
        'rate': 60, # 15 min
    }
    logging.basicConfig(level="INFO", format='%(asctime)s - %(name)s - %(message)s')
    e = DarkSkyDriver(cfg)
    #e = DarkSkyPredictionDriver(cfg)
    e.begin()
