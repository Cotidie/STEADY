#include <string>
#include <vector>
#include <map>
#include <iostream>

using namespace std;

vector<string> split(string s, string delim) {
    vector<string> result;
    int pos = 0;
    string token;
    
    while (pos = s.find(delim), pos != string::npos ) {
        string token = s.substr(0, pos);
        result.push_back(token);
        s.erase(0, pos+delim.length());
    }
    result.push_back(s);
    
    return result;
}

// 1. 차량 번호별 총 주차 시간을 구한다.
// 2. 주차 요금을 계산한다.
class Time {
public:
    Time() {}
    Time(int minutes): minutes(minutes) {}
    Time(string time) {
        int mid = time.find(":");
        int hour = stoi(time.substr(0, 2));
        int minute = stoi(time.substr(mid+1, 2));
        
        this->minutes = hour * 60 + minute;
    }
    
    Time operator-(const Time& other) { return Time(minutes - other.minutes); }

    int minutes;
};

struct Ticket {
    Time timeIn;
    bool expired;
};

class ParkingLot {
public:
    ParkingLot(vector<int>& fees) {
        baseMeter = fees[0]; baseCost = fees[1];
        unitMeter = fees[2]; unitCost = fees[3];
    }  
    
    void parkIn(string slate, Time time) {
        tickets[slate] = Ticket{time, false};
        if (meter.find(slate) == meter.end()) meter[slate] = 0;
    }
    
    void parkOut(string slate, Time time) {
        int minutes = (time - tickets[slate].timeIn).minutes;
        meter[slate] += minutes;
        tickets[slate].expired = true;
    }
    
    void close() {
        for (auto pair : tickets) {
            string plate = pair.first;
            Ticket& ticket = pair.second;
            if (ticket.expired) continue;
            
            parkOut(plate, Time("23:59"));
        }
    }
    
    vector<int> getCosts() {
        vector<int> costs;
        
        for (auto pair : meter) {
            int minutes = pair.second;
            int cost = calculateCost(minutes);
            
            costs.push_back(cost);
        }
        
        return costs;
    }
    
private:
    int calculateCost(int meter) {
        if (meter <= baseMeter) return baseCost;
        
        meter -= baseMeter;
        int times = meter / unitMeter;
        if (meter % unitMeter > 0) times++;
        
        return baseCost + unitCost*times;
    }
    
    int baseMeter, unitMeter;
    int baseCost, unitCost;
    
    map<string, int> meter;
    map<string, Ticket> tickets;
};

vector<int> solution(vector<int> fees, vector<string> records) {
    ParkingLot lot(fees);
    
    for (string& record : records) {
        vector<string> splitted = split(record, " ");
        
        Time time = Time(splitted[0]);
        string slate = splitted[1];
        string action = splitted[2];
        
        if (action == "IN") lot.parkIn(slate, time);
        if (action == "OUT") lot.parkOut(slate, time);
    }
    
    lot.close();
    
    return lot.getCosts();
}