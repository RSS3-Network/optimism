import json


def merge():
    # read all lines from csb-state-dump.json
    csb_state = []
    with open('./csb-state-dump.json', 'r') as file:
        for line in file:
            json_data = json.loads(line)
            csb_state.append(json_data)

    # read from l2 genesis.json
    with open('./genesis.json', 'r') as file:
        l2_genesis = json.load(file)

    # parse csb_state and merge states into l2_genesis
    total = 0
    for item in csb_state:
        address = item.pop('address', None)
        if address:
            # skip empty account
            if 0 == item['nonce'] and item['codeHash'] == '0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470':
                print("skip", address)
                continue

            # delete unused items
            item.pop("root", None)
            item.pop("key", None)
            item.pop("codeHash", None)
            # clear csb balance
            item['balance'] = '0x0'
            # merge
            l2_genesis['alloc'][address] = item

            total = total + 1

    # save
    with open('merged_genesis.json', 'w') as file:
        json.dump(l2_genesis, file, indent=4)

    print("save merged file to merged_genesis.json, total:", total)


def main():
    merge()


if __name__ == '__main__':
    main()
