import { StatusBar } from 'expo-status-bar';
import { FC, useState } from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { GetImagesButton } from './src/components/Widget';
import { theme } from './src/theme';

import * as fs from 'expo-file-system';

const App: FC = () => {
    const [img, setImg] = useState<string>('Valor inicial');

    const onPress = () => {
        const docDir = fs.documentDirectory as string;
        const localFile = `${docDir}`;

        console.log('### localFile: ', localFile);

        fs.StorageAccessFramework.readDirectoryAsync(localFile)
            .then((data) => {
                console.log('\n\n### data: ', data);
                console.log('\n\n### data length: ', data.length);

                data.forEach((item, index) => {
                    console.log(`\n\nindex: ${index}`);
                    console.log(item);
                });

                setImg;
                // const base64 = 'data:image/jpg;base64' + data;
                // setImg(base64)
            })
            .catch((err) => {
                console.log('### getImgError: ', err);
            });
    };

    return (
        <View style={styles.container}>
            <StatusBar style="light" backgroundColor="transparent" translucent />

            <Text style={{ color: '#fff' }}>{img}</Text>

            <GetImagesButton onPress={onPress} />
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: theme.colors.background,
        alignItems: 'center',
        justifyContent: 'center',
    },
});

export default App;
