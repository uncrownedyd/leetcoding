#include <iostream>
#include <string>
#include <vector>
#include <unordered_map>
#include <queue>
#include "base_stl/vector_test.h"
#include <ctime>
#include "base_stl/bigInt.h"
using namespace std;
using vvb = vector<vector<bool>>;
using vvi = vector<vector<int>>;

class Solution {
public:

    void initPq(vvi &heightMap, int w, int l, vvb &visited) {
        for (int i = 0; i < w; ++i) {
            for (int j = 0; j < l; ++j) {
                if (i == 0 || i == w - 1 || j == 0 || j == l - 1) {
                    q.push({i, j, heightMap[i][j]});
                    visited[i][j] = true;
                }
            }
        }
    }
    int trapRainWater(vector<vector<int>>& heightMap) {
        if (heightMap.size() <= 2 || heightMap[0].size() <= 2) {
            return 0;
        }
        int w = heightMap.size();
        int l = heightMap[0].size();
        vvb visited(w, vector<bool>(l, false));
        initPq(heightMap, w, l, visited);

        int ans = 0;
        int neibor[4][2] = {{0,1}, {0, -1}, {1, 0}, {-1, 0}};
        while (!q.empty()) {
            BevPoint cur = q.top();
            q.pop();
            for (int k = 0; k < 4; ++k) {
                int x = cur.pX + neibor[k][0];
                int y = cur.pY + neibor[k][1];
                if( x >= 0 && x < w && y >= 0 && y < l && !visited[x][y]) {
                    if (heightMap[x][y] < cur.height) {
                        ans += cur.height - heightMap[x][y];
                    }
                    visited[x][y] = true;
                    int newHeight = max(heightMap[x][y], cur.height);
                    q.push({x, y, newHeight});
                }
            }
        }
        return ans;
    }

private:
    struct BevPoint
    {
        int pX = 0;
        int pY = 0;
        int height = 0;
        BevPoint(int a, int b, int c): pX(a), pY(b), height(c){}
    };
    struct cmp{
        bool operator()(const BevPoint &a, const BevPoint &b) {
            return a.height > b.height;
        }
    };
    priority_queue<BevPoint, vector<BevPoint>, cmp> q;
};

bool Test(Solution &s, vvi &input, const int expect) {
    int ans = s.trapRainWater(input);
    return ans == expect;
}

int main()
{
    Solution s;
    vvi in;
    int expect;

    in = {{1,4,3,1,3,2}, {3,2,1,3,2,4}, {2,3,3,2,3,1}};
    expect = 4;

    in = {{3,3,3,3,3}, {3,2,2,2,3}, {3,2,1,2,3}, {3,2,2,2,3}, {3,3,3,3,3}};
    expect = 10;

    cout<<Test(s, in, expect);
    return 0;
}