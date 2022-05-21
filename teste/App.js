import { StatusBar } from 'expo-status-bar';
import { useState } from 'react';
import { Button, StyleSheet, Text, View } from 'react-native';

// import * as fs from 'react-native-fs'

export default function App() {
  const [img, setImg] = useState('Valor inicial');

  const onProcess = () => {
    alert('teste');

    setImg;
    // console.log('### storage: ', fs.PicturesDirectoryPath);
  }

  return (
    <View style={styles.container}>
      <StatusBar style="light" translucent />

      <View style={styles.card}>
        <Text style={styles.title}>Img</Text>

        <Text style={styles.text}>{img}</Text>
      </View>

      <View style={styles.button}>
        <Button title='tirar print' onPress={onProcess} />
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#333',
    alignItems: 'center',
    justifyContent: 'center',
  },
  card: {
    minWidth: '90%',
    minHeight: 200,
    marginBottom: 100,
    padding: 20,
    borderRadius: 10,
    color: '#000',
    backgroundColor: '#fff',
  },
  title: {
    color: '#000',
    fontWeight: 'bold',
    fontSize: 20,
    marginBottom: 20,
  },
  text: {
    color: '#000',
  },
  button: {
    width: '90%',
    top: 0,
    bottom: 0,
    backgroundColor: '#fff',
  }
});
