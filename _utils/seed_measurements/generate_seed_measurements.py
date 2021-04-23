import os
import requests
from datetime import datetime, timedelta
import random
import psycopg2
import psycopg2.extras

########################################################################
def db_connection():
    
    return psycopg2.connect(
        user=os.getenv('INSTRUMENTATION_DBUSER', default='x'),
        host=os.getenv('INSTRUMENTATION_DBHOST', default='x'),
        dbname=os.getenv('INSTRUMENTATION_DBNAME', default='x'),
        password=os.getenv('INSTRUMENTATION_DBPASS', default='x')
    )
########################################################################
def generate_measurements(timeseries_id, param, measurements_sql):

    # This could defintely be done better using another library like
    # pandas.  Feel free to rewrite it to get better sample data

    value_count = 30
    dt = datetime.utcnow().replace(minute=0)

    if param.lower() == 'stage':
        samplelist = random.sample(range(10, 60), value_count)
    elif param.lower() == 'elevation':
        samplelist = random.sample(range(550, 850), value_count)
    elif param.lower() == 'voltage':
        samplelist = list(range(9,14))
    elif param.lower() == 'water-temperature':
        samplelist = random.sample(range(30, 80), value_count)
    elif param.lower() == 'precipitation':
        samplelist = list(range(20,35))[::-1]
    elif param.lower() == 'ph':
        samplelist = list(range(2,12))
    else:
        samplelist = random.sample(range(50, 80), value_count)
   
    # extend list with same list but in reverse (mirror image)
    if param.lower() != 'precipitation':
        samplelist.extend(samplelist[::-1])
    
    for x in samplelist:
        measurements_sql += f"('{timeseries_id}', '{dt.strftime('%Y-%m-%dT%H:%M:%SZ')}', {x}),\n"
        dt = dt - timedelta(hours=6)

    return measurements_sql
########################################################################

API_HOST = f"http://{os.getenv('INSTRUMENTATION_API_HOST', default='http://instrumentation-api_api_1')}"
INSTRUMENTATION_ROUTE_PREFIX = os.getenv('INSTRUMENTATION_ROUTE_PREFIX', default='')

script_dir = os.path.dirname(os.path.realpath(__file__))
measurements_sql = 'INSERT INTO timeseries_measurement (timeseries_id, time, value) VALUES\n'

print('Fetching timeseries from Instrumentation API')
r = requests.get(f'{API_HOST}{INSTRUMENTATION_ROUTE_PREFIX}/timeseries')
timeseries = r.json()


for ts in timeseries:
    r = requests.get(f"{API_HOST}{INSTRUMENTATION_ROUTE_PREFIX}/timeseries/{ts['id']}/measurements")
    ts_measurements = r.json()
    # Only add sample data if none exists
    if ts_measurements['items'] is None:
        measurements_sql = generate_measurements(ts['id'], ts['name'], measurements_sql)

# remove the last comma and new line and replace with ; and \n
measurements_sql = measurements_sql[:-2]+';\n'

print('Inserting sample data measurements into DB...')
# Connect to database and insert payload seed data
try:
    conn = db_connection()
    c = conn.cursor()
    c.execute(measurements_sql)
    conn.commit()
except Exception as e:
    print(e)
finally:
    c.close()
    conn.close()

# with open(f'{script_dir}/seed_measurements.sql', 'w+') as f:
#     f.write(measurements_sql)

print('################################')
print('Sample measurements data loaded.')
print('################################')
exit(0)