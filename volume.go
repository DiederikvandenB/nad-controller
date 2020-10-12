package main

func getDecibels(percentage string) string {
  decibels := map[string]string {
    "0": "-80.0",
    "0.01": "-79.5",
    "0.02": "-79.0",
    "0.03": "-77.5",
    "0.04": "-76.5",
    "0.05": "-75.0",
    "0.06": "-73.5",
    "0.07": "-72.5",
    "0.08": "-71.0",
    "0.09": "-69.5",
    "0.1": "-68.5",
    "0.11": "-67.0",
    "0.12": "-66.0",
    "0.13": "-64.5",
    "0.14": "-63.0",
    "0.15": "-62.0",
    "0.16": "-60.5",
    "0.17": "-59.0",
    "0.18": "-58.0",
    "0.19": "-56.5",
    "0.2": "-55.5",
    "0.21": "-54.5",
    "0.22": "-53.5",
    "0.23": "-52.5",
    "0.24": "-51.5",
    "0.25": "-50.5",
    "0.26": "-49.5",
    "0.27": "-48.5",
    "0.28": "-47.5",
    "0.29": "-46.5",
    "0.3": "-45.5",
    "0.31": "-44.5",
    "0.32": "-43.5",
    "0.33": "-42.5",
    "0.34": "-41.5",
    "0.35": "-40.5",
    "0.36": "-39.5",
    "0.37": "-38.5",
    "0.38": "-37.5",
    "0.39": "-36.5",
    "0.4": "-35.5",
    "0.41": "-34.5",
    "0.42": "-33.5",
    "0.43": "-32.5",
    "0.44": "-31.5",
    "0.45": "-30.5",
    "0.46": "-29.5",
    "0.47": "-28.5",
    "0.48": "-27.5",
    "0.49": "-26.5",
    "0.5": "-25.5",
    "0.51": "-24.5",
    "0.52": "-23.5",
    "0.53": "-22.5",
    "0.54": "-21.5",
    "0.55": "-20.5",
    "0.56": "-19.5",
    "0.57": "-18.5",
    "0.58": "-17.5",
    "0.59": "-16.5",
    "0.6": "-15.5",
    "0.61": "-14.5",
    "0.62": "-14.0",
    "0.63": "-13.0",
    "0.64": "-12.5",
    "0.65": "-11.5",
    "0.66": "-11.0",
    "0.67": "-10.0",
    "0.68": "-9.5",
    "0.69": "-8.5",
    "0.7": "-8.0",
    "0.71": "-7.0",
    "0.72": "-6.5",
    "0.73": "-5.5",
    "0.74": "-5.0",
    "0.75": "-4.0",
    "0.76": "-3.5",
    "0.77": "-2.5",
    "0.78": "-2.0",
    "0.79": "-1.0",
    "0.8": "0",
    "0.81": "1.0",
    "0.82": "1.5",
    "0.83": "2.0",
    "0.84": "2.5",
    "0.85": "3.5",
    "0.86": "4.0",
    "0.87": "4.5",
    "0.88": "5.0",
    "0.89": "5.5",
    "0.9": "6.5",
    "0.91": "7.0",
    "0.92": "7.5",
    "0.93": "8.0",
    "0.94": "8.5",
    "0.95": "9.5",
    "0.96": "10.0",
    "0.97": "10.5",
    "0.98": "11.0",
    "0.99": "11.5",
    "1": "12.0",
  }

  return decibels[percentage]
}

func getPercentages(decibel string) string {
  percentages := map[string]string {
    "-80.0": "0",
    "-79.5": "0.01",
    "-79.0": "0.02",
    "-78.5": "0.02",
    "-78.0": "0.02",
    "-77.5": "0.03",
    "-77.0": "0.03",
    "-76.5": "0.04",
    "-76.0": "0.04",
    "-75.5": "0.04",
    "-75.0": "0.05",
    "-74.5": "0.05",
    "-74.0": "0.05",
    "-73.5": "0.06",
    "-73.0": "0.06",
    "-72.5": "0.07",
    "-72.0": "0.07",
    "-71.5": "0.07",
    "-71.0": "0.08",
    "-70.5": "0.08",
    "-70.0": "0.08",
    "-69.5": "0.09",
    "-69.0": "0.09",
    "-68.5": "0.10",
    "-68.0": "0.10",
    "-67.5": "0.10",
    "-67.0": "0.11",
    "-66.5": "0.11",
    "-66.0": "0.12",
    "-65.5": "0.12",
    "-65.0": "0.12",
    "-64.5": "0.13",
    "-64.0": "0.13",
    "-63.5": "0.13",
    "-63.0": "0.14",
    "-62.5": "0.14",
    "-62.0": "0.15",
    "-61.5": "0.15",
    "-61.0": "0.15",
    "-60.5": "0.16",
    "-60.0": "0.16",
    "-59.5": "0.16",
    "-59.0": "0.17",
    "-58.5": "0.17",
    "-58.0": "0.18",
    "-57.5": "0.18",
    "-57.0": "0.18",
    "-56.5": "0.19",
    "-56.0": "0.19",
    "-55.5": "0.20",
    "-55.0": "0.20",
    "-54.5": "0.21",
    "-54.0": "0.21",
    "-53.5": "0.22",
    "-53.0": "0.22",
    "-52.5": "0.23",
    "-52.0": "0.23",
    "-51.5": "0.24",
    "-51.0": "0.24",
    "-50.5": "0.25",
    "-50.0": "0.25",
    "-49.5": "0.26",
    "-49.0": "0.26",
    "-48.5": "0.27",
    "-48.0": "0.27",
    "-47.5": "0.28",
    "-47.0": "0.28",
    "-46.5": "0.29",
    "-46.0": "0.29",
    "-45.5": "0.30",
    "-45.0": "0.30",
    "-44.5": "0.31",
    "-44.0": "0.31",
    "-43.5": "0.32",
    "-43.0": "0.32",
    "-42.5": "0.33",
    "-42.0": "0.33",
    "-41.5": "0.34",
    "-41.0": "0.34",
    "-40.5": "0.35",
    "-40.0": "0.35",
    "-39.5": "0.36",
    "-39.0": "0.36",
    "-38.5": "0.37",
    "-38.0": "0.37",
    "-37.5": "0.38",
    "-37.0": "0.38",
    "-36.5": "0.39",
    "-36.0": "0.39",
    "-35.5": "0.40",
    "-35.0": "0.40",
    "-34.5": "0.41",
    "-34.0": "0.41",
    "-33.5": "0.42",
    "-33.0": "0.42",
    "-32.5": "0.43",
    "-32.0": "0.43",
    "-31.5": "0.44",
    "-31.0": "0.44",
    "-30.5": "0.45",
    "-30.0": "0.45",
    "-29.5": "0.46",
    "-29.0": "0.46",
    "-28.5": "0.47",
    "-28.0": "0.47",
    "-27.5": "0.48",
    "-27.0": "0.48",
    "-26.5": "0.49",
    "-26.0": "0.49",
    "-25.5": "0.50",
    "-25.0": "0.50",
    "-24.5": "0.51",
    "-24.0": "0.51",
    "-23.5": "0.52",
    "-23.0": "0.52",
    "-22.5": "0.53",
    "-22.0": "0.53",
    "-21.5": "0.54",
    "-21.0": "0.54",
    "-20.5": "0.55",
    "-20.0": "0.55",
    "-19.5": "0.56",
    "-19.0": "0.56",
    "-18.5": "0.57",
    "-18.0": "0.57",
    "-17.5": "0.58",
    "-17.0": "0.58",
    "-16.5": "0.59",
    "-16.0": "0.59",
    "-15.5": "0.60",
    "-15.0": "0.60",
    "-14.5": "0.61",
    "-14.0": "0.62",
    "-13.5": "0.62",
    "-13.0": "0.63",
    "-12.5": "0.64",
    "-12.0": "0.64",
    "-11.5": "0.65",
    "-11.0": "0.66",
    "-10.5": "0.66",
    "-10.0": "0.67",
    "-9.5": "0.68",
    "-9.0": "0.68",
    "-8.5": "0.69",
    "-8.0": "0.70",
    "-7.5": "0.70",
    "-7.0": "0.71",
    "-6.5": "0.72",
    "-6.0": "0.72",
    "-5.5": "0.73",
    "-5.0": "0.74",
    "-4.5": "0.74",
    "-4.0": "0.75",
    "-3.5": "0.76",
    "-3.0": "0.76",
    "-2.5": "0.77",
    "-2.0": "0.78",
    "-1.5": "0.78",
    "-1.0": "0.79",
    "-0.5": "0.79",
    "0": "0.80",
    "0.5": "0.80",
    "1.0": "0.81",
    "1.5": "0.82",
    "2.0": "0.83",
    "2.5": "0.84",
    "3.0": "0.84",
    "3.5": "0.85",
    "4.0": "0.86",
    "4.5": "0.87",
    "5.0": "0.88",
    "5.5": "0.89",
    "6.0": "0.89",
    "6.5": "0.90",
    "7.0": "0.91",
    "7.5": "0.92",
    "8.0": "0.93",
    "8.5": "0.94",
    "9.0": "0.94",
    "9.5": "0.95",
    "10.0": "0.96",
    "10.5": "0.97",
    "11.0": "0.98",
    "11.5": "0.99",
    "12.0": "0.10",
  }

  return percentages[decibel]
}