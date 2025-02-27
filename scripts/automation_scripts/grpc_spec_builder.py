import subprocess
import os

TEMPLATE = """
    ,{
          "name": "###METHOD###",
          ###RESTPART###
    }
"""

# must have rest api already implemented

# Lava:
# grpc_server = "public-rpc.lavanet.xyz:9090"
# spec_current_file_path = "/home/user/go/src/lava/cookbook/spec_add_lava.json"

# Osmosis:
# grpc_server = "grpc.osmosis.zone:9090"
# spec_current_file_path = "/home/user/go/src/lava/cookbook/spec_add_osmosis.json" 

# Cosmos
# grpc_server = "gaia-node-1.lavapro.xyz:9090"
# spec_current_file_path = "/home/user/go/src/lava/cookbook/spec_add_cosmoshub.json" 

grpc_server = "juno-node-1.lavapro.xyz:9090"
spec_current_file_path = "/home/user/go/src/lava/cookbook/spec_add_juno.json" 

# 
# grpcurl -plaintext prod-pnet-osmosisnode-1.lavapro.xyz:9090 list # COS3
# grpcurl -plaintext public-rpc.lavanet.xyz:9090 list # LAV1
# grpcurl -plaintext gaia-node-1.lavapro.xyz:9090 list # COS5
# grpcurl -plaintext juno-node-1.lavapro.xyz:9090 list # JUNO
# # --- #
# grpcurl -plaintext gaia-node-1.lavapro.xyz:9092 list # COS5T
# grpcurl -plaintext juno-node-1.lavapro.xyz:9092 list # JUNO-TEST


special_cases_descriptors_with_no_rest_api = []

res = str(subprocess.check_output(f"grpcurl -plaintext {grpc_server} list",shell=True))[2:-1]
arr = res.split("\\n")
spec_res = ""
with open(spec_current_file_path, 'r') as f:
    spec_data = f.read()
    if spec_data.count("chainid") > 1: # remove the 2nd chainid in one spec.
        print("splitting")
        spec_data = spec_data.rsplit("chainid",1)[0]

all_rest_apis = spec_data.split('"name": "/')[1:]
rest_api_list = []
for rest_api in all_rest_apis:
    rest_api = rest_api.split("\"",1)[0]
    rest_api_list.append("/"+rest_api)

original_api_list = rest_api_list[:]
total_number_of_descriptors = 0
rest_lines_not_in_spec = []

skip_services = ["tendermint.liquidity.v1beta1.Query","testdata","reflection"]

def check_skip(service):
    # print("checking serice skip on: %s" % service)
    if service == "":
        return True
    for s in skip_services:
        if s in service:
            return True
    return False        

for service in arr:
    service = service.strip()
    if check_skip(service):
        continue
    print("Processing: ",service)
    descriptors = str(subprocess.check_output(f"grpcurl -plaintext {grpc_server} describe {service}",shell=True))[2:-1]
    print(descriptors)
    descriptors_arr = descriptors.split("rpc ")[1:]
    for descriptor in descriptors_arr:
        total_number_of_descriptors+=1
        if descriptor.strip() == "": 
            continue
        method_name = descriptor.split(" ")[0]
        full_name = f"{service}.{method_name}"
        print("Full name: ", full_name)
        if "get:" not in descriptor:
            if "post:" not in descriptor:
                special_cases_descriptors_with_no_rest_api.append("service: " + service + "\n descriptor: "+descriptor)
                continue
            rest_line = descriptor.split("post:",1)[-1]
            rest_line = rest_line.strip().split("\"",1)[-1]
            rest_line = rest_line.split("\"",1)[0].strip()
        else:
            rest_line = descriptor.split("get:",1)[-1]
            rest_line = rest_line.strip().split("\"",1)[-1]
            rest_line = rest_line.split("\"",1)[0].strip()
       
        print("Rest line: ", rest_line,"\nDescriptor: \n",descriptor)
        # fetch the part from spec:
        if rest_line not in rest_api_list:
            # [print(x) for x in original_api_list]
            try: 
                raise ValueError("rest_line not found in rest_api_list", rest_line)
            except:
                pass
        else:
            rest_api_list.remove(rest_line)

        if (rest_line+"\",") not in spec_data:
            rest_lines_not_in_spec.append(rest_line)
            continue
            # raise ValueError("rest_line not in spec:\n" + rest_line)
        spec_body = spec_data.split(rest_line+"\",",1)[-1].split('"name":',1)[0].rsplit("},",1)[0].strip()
        spec_body = spec_body.replace('"interface": "rest",','"interface": "grpc",',2)
        if '"interface": "rest"' in spec_body:
            raise ValueError("parsing error", "\nDescriptor\n",descriptor,"\nFullname\n",full_name )
        final_part = TEMPLATE.replace("###METHOD###",full_name).replace("###RESTPART###",spec_body).strip()
        print("final_part",final_part)
        spec_res += final_part

        print("@@@@@@@@@@@@@@\n\n")
        # print(spec_res)

print("total number of api's added:", total_number_of_descriptors)
print("rest apis that werent added from rest:")
[print(x) for x in rest_api_list]
print("@@@@@@@@@@@@@@\n\n")
if len(rest_lines_not_in_spec) > 0:
    [print(x) for x in rest_lines_not_in_spec]
    raise ValueError("not all lines were found in spec, missing lines")
if len(special_cases_descriptors_with_no_rest_api) > 0:
    print("Special cases rpc:")
    [print(x) for x in special_cases_descriptors_with_no_rest_api]

with open("/home/user/go/src/lava/scripts/automation_scripts/automation_results/spec_add.json", "w+") as json_file:
    json_file.write(spec_res)


    

