from flask import Flask, jsonify, request

app = Flask(__name__)

receitas = [
    {
        'id': 1,
        'nome': 'Bolo de Chocolate Simples',
        'ingredientes': [
            '2 xícaras de farinha',
            '1 xícara de chocolate em pó',
            '3 ovos',
            '1 xícara de leite',
            '1/2 xícara de óleo',
            '1 xícara de açúcar',
            '1 colher de sopa de fermento'
        ],
        'instrucoes': '1. Misture os ingredientes secos. 2. Adicione os líquidos e misture bem. 3. Leve ao forno por 40 minutos.'
    },
    {
        'id': 2,
        'nome': 'Salada de Quinoa com Vegetais',
        'ingredientes': [
            '1 xícara de quinoa',
            '2 xícaras de água',
            '1 pepino picado',
            '1 tomate picado',
            '1/2 cebola roxa picada',
            'Suco de 1 limão'
        ],
        'instrucoes': '1. Cozinhe a quinoa. 2. Deixe esfriar. 3. Misture com os vegetais e tempere.'
    },
    {
        'id': 3,
        'nome': 'Smoothie de Frutas',
        'ingredientes': [
            '1 banana congelada',
            '1/2 xícara de morangos',
            '1/2 xícara de iogurte natural',
            '1/4 xícara de leite'
        ],
        'instrucoes': '1. Coloque todos os ingredientes no liquidificador. 2. Bata até ficar cremoso.'
    }
]

@app.route('/receitas', methods=['GET'])
def get_receitas():
    return jsonify(receitas)

@app.route ('/receitas/<int:receita_id>', methods=['GET'])
def get_receita(receita_id):
    receita = next((r for r in receitas if r['id'] == receita_id), None)
    if receita:
        return jsonify(receita)
    else:
        return jsonify({'error': 'Receita não encontrada'}), 404

@app.route ('/receitas/', methods=['POST'])
def add_receita():
    nova_receita = request.get_json()

    if nova_receita or 'nome' not in nova_receita or 'ingredientes' not in nova_receita:
        return jsonify({'error': 'Dados inválidos'}), 400
    
    if receitas:
        nova_receita['id'] = receitas[-1]['id'] + 1
    else:
        nova_receita['id'] = 1

    receitas.append(nova_receita)
    return jsonify(nova_receita), 201

@app.route ('/receitas/<int:receita_id>', methods=['DELETE'])
def delete_receita(receita_id):
    global receitas
    receitas = [r for r in receitas if r['id'] != receita_id]
    return jsonify({'message': 'Receita deletada'}), 200


@app.route ('/receitas/<int:receita_id>', methods=['PUT'])
def update_receita(receita_id): 
    receita_atualizada = request.get_json()
    receita = next((r for r in receitas if r['id'] == receita_id), None)
    
    if not receita:
        return jsonify({'error': 'Receita não encontrada'}), 404
    
    if not receita_atualizada or 'nome' not in receita_atualizada or 'ingredientes' not in receita_atualizada:
        return jsonify({'error': 'Dados inválidos'}), 400
    
    receita.update(receita_atualizada)
    return jsonify(receita), 200

if __name__ == '__main__':
    app.run(debug=True)
