# Atividade Avaliativa 4 - Respostas

Nome: Eduardo F. Pressutto e Lucas W. F. Pepes

GRR: 20252190 e 20255692

Se a atividade foi feita em dupla, escreva aqui os dois nomes e os dois GRR. As
duas pessoas enviam o mesmo `.zip` na UFPR Virtual.

## Pergunta

1. Explique por que sua implementação respeita a complexidade computacional exigida
   pelo enunciado: preprocessamento do grafo em tempo linear e apresentação da
   distância entre qualquer par de vértices em tempo constante.

   Pois aproveita que o grafo é uma linha, possuindo [ v - 1 ]  arestas, e cada vértice tendo grau no máximo 2, no preprocessamento é feito primeiro uma busca pelo primeiro vertice de grau 1, que no pior dos casos percorre todos os vertices custando O(V). Após isso o algoritmo percorre a linha a partir desse extremo, visitando cada vértice apenas uma vez, e como cada vértice possui no maximo duas arestas, a quantidade de operações realizadas em cada iteração é constante, fazendo com que essa etapa também custe O(v).
   Durante esse fluxo, é armazenado em distExtremo[v] a distancia acumulada entre cada vértice e o extremo escolhido da linha. O vetor possui uma posição por vértice, logo o espaço adicional utilizado é O(V). Para responder uma consulta Distancia() acessa duas posições correspondentes aos vértices, calcula a diferença entre as distâncias armazenadas e aplica o valor. Como as operações realizadas são de quantidade fixa, não importa o numero de vertices, cada consulta custa O(1). E isso se justifica pois em uma linha a distancia entre dois vértices é igual a diferença absoluta entre suas posições acumuladas a partir do mesmo extremo.

## Uso de IA

Diga se usou ferramentas de IA generativa nesta atividade e em que partes. Não
influencia a nota.

Não foi utilizado IA generativa nessa atividade
