import json,random
from datetime import datetime
#https://github.com/rpietro/GlocalRegistry/blob/master/MINEdata.csv -> Reference
Temperature_Data = {}
record_data = {}
Insulin_Data = {}
BO_Data = {}
BR_Data = {}
BP_Data = {}
BA_Data = {}
i=0
with open('sensorPacemaker.json','w',encoding='utf-8') as fjson:
    while True:
        record_data['atrial_heart_rate_basal'] = random.randint(60,90)
        record_data['ventric_heart_rate_basal'] = random.randint(30,50)
        record_data['qrs_duration_basal'] = random.randint(90,150)
        record_data['hemoglobin_basal'] = float("{0:.2f}".format(random.uniform(10, 17)))
        record_data['platelets_basal'] = random.randint(165000,240000)
        record_data['leukocytes_basal'] = random.randint(6500,10000)
        record_data['sodium_basal'] = random.randint(130,150)
        record_data['urea_basal'] = random.randint(40,80)
        record_data['total_procedure_time'] = random.randint(60,140)
        if record_data['atrial_heart_rate_basal'] > 75 and record_data['platelets_basal'] > 200000 :
            record_data['Condition'] = "Normal"
        else:
            record_data['Condition'] = "Abnormal"
        i += 1
        record = json.dumps(record_data)
        print('Print Data',record_data)
        if i == 100:
            break
        fjson.write('{}\n'.format(record))
    fjson.close() 

j = 0
with open('sensorTemprature.json','a',encoding='utf-8') as fjson:
    while True:
        Temperature_Fake_Value = float("{0:.2f}".format(random.uniform(97, 108)))
        Temperature_Data['Date'] = (datetime.today()).strftime("%d-%b-%Y %H:%M:%S:%f")
        Temperature_Data['Temperature'] = Temperature_Fake_Value
        if Temperature_Fake_Value > 100:
            Temperature_Data['Condition'] = "Abnormal"
        else:
            Temperature_Data['Condition'] = "Normal"
            
        j += 1
        record = json.dumps(Temperature_Data)
        print('Print Data',Temperature_Data)
        if j == 100:
            break
        fjson.write('{}\n'.format(record))
    fjson.close()
    
k = 0
with open('sensorInsulin.json','a',encoding='utf-8') as fjson:
    while True:
        Insulin_Data['Date'] = (datetime.today()).strftime("%d-%b-%Y %H:%M:%S:%f")
        Insulin_Data['Glucose Specimen'] = random.randint(100,118)
        if Insulin_Data['Glucose Specimen'] > 110:
            Insulin_Data['Condition'] = "Abnormal"
        else:
            Insulin_Data['Condition'] = "Normal"
        
        Insulin_Data['Insulin Specimen'] = random.randint(12,22)
        if Insulin_Data['Insulin Specimen'] > 17:
            Insulin_Data['Condition'] = "Abnormal"
        else:
            Insulin_Data['Condition'] = "Normal"
            
        k += 1
        record = json.dumps(Insulin_Data)
        print('Print Data',Insulin_Data)
        if k == 100:
            break
        fjson.write('{}\n'.format(record))
    fjson.close()

p = 0
with open('sensorBodyOxygen.json','a',encoding='utf-8') as fjson:
    while True:
        BO_Data['Date'] = (datetime.today()).strftime("%d-%b-%Y %H:%M:%S:%f")
        BO_Data['BO Specimen'] = random.randint(88,105)
        if BO_Data['BO Specimen'] > 95:
            BO_Data['Condition'] = "Normal"
        else:
            BO_Data['Condition'] = "Abnormal"
        p += 1
        record = json.dumps(BO_Data)
        print('Print Data',BO_Data)
        if p == 100:
            break
        fjson.write('{}\n'.format(record))
    fjson.close()

q = 0
with open('sensorBreathingRate.json','a',encoding='utf-8') as fjson:
    while True:
        BR_Data['Date'] = (datetime.today()).strftime("%d-%b-%Y %H:%M:%S:%f")
        BR_Data['Breathing Rate'] = random.randint(5,30)
        if BR_Data['Breathing Rate'] < 12 and BR_Data['Breathing Rate'] > 20 :
            BR_Data['Condition'] = "Abnormal"
        else:
            BR_Data['Condition'] = "Normal"
        q += 1
        record = json.dumps(BR_Data)
        print('Print Data',BR_Data)
        if q == 100:
            break
        fjson.write('{}\n'.format(record))
    fjson.close()
    
r = 0
with open('sensorBloodPressure.json','a',encoding='utf-8') as fjson:
    while True:
        BP_Data['Date'] = (datetime.today()).strftime("%d-%b-%Y %H:%M:%S:%f")
        BP_Data['Blood Pressure'] = random.randint(5,30)
        if BP_Data['Blood Pressure'] < 12 and BP_Data['Blood Pressure'] > 20 :
            BP_Data['Condition'] = "Abnormal"
        else:
            BP_Data['Condition'] = "Normal"
        r += 1
        record = json.dumps(BP_Data)
        print('Print Data',BP_Data)
        if r == 100:
            break
        fjson.write('{}\n'.format(record))
    fjson.close()
    
s = 0
with open('sensorBloodAlcohol.json','a',encoding='utf-8') as fjson:
    while True:
        BA_Data['Date'] = (datetime.today()).strftime("%d-%b-%Y %H:%M:%S:%f")
        BA_Data['Blood Alcohol'] = float("{0:.2f}".format(random.uniform(0, 1)))
        if BA_Data['Blood Alcohol'] < 12 and BA_Data['Blood Alcohol'] > 20 :
            BA_Data['Condition'] = "Abnormal"
        else:
            BA_Data['Condition'] = "Normal"
        s += 1
        record = json.dumps(BA_Data)
        print('Print Data',BA_Data)
        if s == 100:
            break
        fjson.write('{}\n'.format(record))
    fjson.close()

        