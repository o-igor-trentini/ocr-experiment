import { ChatTeardropDots } from 'phosphor-react-native';
import React, { FC, useState } from 'react';
import { Text, TouchableOpacity } from 'react-native';
import { theme } from '../../theme';
import { styles } from './styles';

import * as fs from 'expo-file-system';

export const Widget: FC = () => {
    const [imgDownloaded, setImgDownloaded] = useState<string>('');

    const onPress = () => {
        const teste = fs.StorageAccessFramework.getUriForDirectoryInRoot('');

        console.log(teste);

        return;

        // const localFile = fs.StorageAccessFramework.getUriForDirectoryInRoot('/Documents');
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

                // const base64 = 'data:image/jpg;base64' + data;

                // setImgDownloaded(base64);
            })
            .catch((err) => {
                console.log('### getImgError: ', err);
            });
    };

    return (
        <>
            <TouchableOpacity onPress={onPress} style={styles.button}>
                <ChatTeardropDots size={24} color={theme.colors.text_on_brand_color} weight="bold" />

                <Text>DL: {imgDownloaded}</Text>
            </TouchableOpacity>
        </>
    );
};
