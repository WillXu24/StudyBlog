// Defines the entry point for the console application.
#include <iostream>
using namespace std;

int main() {
    //sore the ciphertext
    char a[100];

    // store the decrypted results
    char b[100];

    cout << "Caesar cipher: Please input the cipher text:" << endl;
    cin>> a ;
    // brute-force attack: try all the keys(25 in total) and select
    // the decrypted message with semantic meaning as the plain text
    // and you are done!!
    // the plaintext is "watch out for brutus".
    for(int k = 1; k < 26; k++) {
        cout << "decrypted result #" << k << "(shift parameter):";
        for (int i = 0; i < 99 && a[i] != 0; i++) {
            if (a[i] + k > 'z')
                b[i] = a[i] + k - 26;
            else b[i] = a[i] + k;


            cout << b[i];

        }
        cout << endl;
    }
    cout << endl;
    system("PAUSE");
    return 0;
}

