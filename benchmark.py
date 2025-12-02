import time

def process_log():
    count = 0
    error_count = 0
    
    with open("dummy_logs_milion.txt", "r") as f:
        for line in f:
            count += 1
            if "ERROR" in line:
                error_count += 1
    
    return count, error_count

if __name__ == "__main__":
    start = time.time()
    print("starting python processing...")
    
    total, errors = process_log()
    end = time.time()
    duration = end - start
    
    print(f"Python Processing took: {duration:.4f} seconds")
    print(f"Total Logs: {total}")
    print(f"Total Errors: {errors}")
    