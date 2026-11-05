1. for received := range channel don't require additional receiver
2. worker must contain channel close, else there will be deadlock: worker ends his shift, but receiver still waiting new records.